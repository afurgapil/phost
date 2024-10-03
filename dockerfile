FROM golang:1.22 AS backend-builder

WORKDIR /app

COPY backend/go.mod backend/go.sum ./backend/
RUN cd backend && go mod download

COPY database/go.mod database/go.sum ./database/
RUN cd database && go mod download

COPY backend ./backend
COPY database ./database

RUN cd database && go build -tags netgo -ldflags '-s -w' -o /bin/database cmd/phost/main.go

RUN cd backend && go build -tags netgo -ldflags '-s -w' -o /bin/backend cmd/phost-backend/main.go

FROM node:18 AS frontend-builder

WORKDIR /app

COPY frontend/package.json frontend/package-lock.json ./
RUN npm ci

COPY frontend ./

RUN npm run build


FROM alpine:latest

RUN apk add --no-cache nodejs npm

WORKDIR /root/

COPY --from=backend-builder /bin/database .
COPY --from=backend-builder /bin/backend .

COPY --from=frontend-builder /app/out ./frontend

ENV ENCRYPTION_KEY=thisIsA32ByteLongPassphraseeeee!
ENV BACKEND_PORT=8081
ENV DB_PORT=8080
ENV FRONTEND_PORT=3000

EXPOSE $DB_PORT
EXPOSE $BACKEND_PORT
EXPOSE $FRONTEND_PORT

RUN npm install -g serve

CMD sh -c "./database & \
           ./backend & \
           serve -s frontend -l $FRONTEND_PORT"

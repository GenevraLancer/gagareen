FROM nginx:1.18

COPY nginx/nginx.conf /etc/nginx/nginx.conf
COPY backend/cert/server.crt /etc/nginx/server.crt
COPY backend/cert/server.key /etc/nginx/server.key

CMD ["nginx", "-g", "daemon off;"]
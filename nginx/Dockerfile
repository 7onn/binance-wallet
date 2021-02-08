
FROM nginx:alpine
RUN chmod -R 766 /var/log/nginx && \
  unlink /var/log/nginx/error.log && \
  ln -sf /dev/stderr /var/log/nginx/stderr.log 

COPY ./nginx.conf  /etc/nginx/nginx.conf
CMD ["nginx", "-g", "daemon off;"]

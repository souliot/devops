FROM scratch
WORKDIR /app
ADD devops.zip .
EXPOSE 8080
CMD ["./devops"]
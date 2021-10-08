FROM scratch
WORKDIR /app
ADD devops.tar.gz .
EXPOSE 8080
CMD ["./devops"]
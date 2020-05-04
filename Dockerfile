FROM alpine

# expose port
EXPOSE 8080

# Add executable into image
ADD build/app /

# execute command when docker launches / run
CMD ["/app"]


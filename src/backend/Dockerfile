FROM alpine
RUN apk add --no-cache ca-certificates
COPY backend /backend
EXPOSE 443
ENTRYPOINT [ "/backend" ]
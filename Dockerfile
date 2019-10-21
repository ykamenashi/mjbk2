FROM alpine
RUN wget 'https://ja.osdn.net/frs/redir.php?m=jaist&f=nkf%2F70406%2Fnkf-2.1.5.tar.gz' -O nkf.tgz && \
 tar zxf nkf.tgz && \
 apk add --no-cache --virtual .build-deps \
 make gcc g++ && \
 cd nkf-* && \
 make && make install && \
 apk del .build-deps
COPY service /
CMD  /service
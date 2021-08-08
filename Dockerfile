FROM golang:1.16

WORKDIR $GOPATH/src/github.com/SomeshSunariwal/golang-okta-token-handler

ENV PORT=8080

# # HOST FOR CORS 
# ARG HOST

COPY bin/GraphQL_implementation .

EXPOSE 8080

CMD ["./GraphQL_implementation"]

# docker run -p 8080:8080 --env-file ./env.list someshdokerbox/graphql-boiler-plate

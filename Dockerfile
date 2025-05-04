FROM cosmtrek/air

COPY . goexpert-final-challenge-1

WORKDIR /goexpert-final-challenge-1

COPY . .

CMD ["air"]
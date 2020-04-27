# Node.js

## Logging

- [Using morgan and winston for logging](https://stackoverflow.com/questions/27906551/node-js-logging-use-morgan-and-winston)

## Exit process on critical errors (i.e. can't connect to database)

```javascript
const PORT = process.env.PORT || 3000;

const server = app.listen(
  PORT,
  console.log(`Server running in ${process.env.NODE_ENV} mode on port ${PORT}`)
);

process.on("unhandledRejection", (err, promise) => {
  console.log(`Error: ${err.message}`);
  server.close(() => process.exit(1));
});
```

## Connect to Mongoose

db.js:

```javascript
const mongoose = require("mongoose");

const db = async () => {
  const conn = await mongoose.connect(process.env.MONGO_URI, {
    useNewUrlParser: true,
    useCreateIndex: true,
    useFindAndModify: false,
    useUnifiedTopology: true,
  });

  console.log(`MongoDB connected: ${conn.connection.host}`);
};

module.exports = db;
```

server.js:

```javascript
const db = require(".db");

db();
```

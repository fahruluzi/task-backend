require("express-group-routes");
require("dotenv").config();

const createError = require("http-errors");
const express = require("express");
const cookieParser = require("cookie-parser");
const cors = require("cors");
const logger = require("morgan");
const port = process.env.PORT;

// router
const apiRouter = require("./application/routes");

// middlewares
const generalMiddleware = require("./middlewares/general");

// database
const database = require("./lib/database");
database.connectDB();

const app = express();

// enable cors
app.use(
    cors({
        exposedHeaders: ["Content-Disposition"],
    })
);

app.use(logger('dev'));
app.use(express.json({limit: "100mb"}));
app.use(express.urlencoded({extended: false}));
app.use(cookieParser());

app.use(generalMiddleware);

// Swagger Settings
const swaggerUi = require('swagger-ui-express');
const openApiDocumentation = require('./swagger.json');
app.use('/swagger', swaggerUi.serve, swaggerUi.setup(openApiDocumentation));

apiRouter(app);

// catch 404 and forward to error controller
app.use(function (req, res, next) {
    next(createError(404));
});

// error controller
app.use(function (err, req, res, next) {
    res.sendError(err, err.message, err.status);
});

//server listening
app.listen(port, () => {
    console.log(`Server is running at port ${port}`);
});

module.exports = app;
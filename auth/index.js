require("express-group-routes");

const createError = require("http-errors");
const express = require("express");
const cookieParser = require("cookie-parser");
const cors = require("cors");
const port = 3000;

// router
const apiRouter = require("./application/routes");

// middlewares
const generalMiddleware = require("./middlewares/general");

const app = express();

// enable cors
app.use(
    cors({
        exposedHeaders: ["Content-Disposition"],
    })
);

app.use(express.json({ limit: "100mb" }));
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());

app.use(generalMiddleware);

// Swagger Settings
const swaggerUi = require('swagger-ui-express');
const openApiDocumentation = require('./swagger.json');
app.use('/swagger', swaggerUi.serve, swaggerUi.setup(openApiDocumentation));

app.group("/api", (router) => {
    apiRouter(router);
});

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
const db = require("diskdb")

module.exports = {
    connectDB: function () {
        db.connect("../data", ["users"])
    }
}
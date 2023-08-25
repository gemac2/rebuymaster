require("expose-loader?exposes=$,jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
require("@fortawesome/fontawesome-free/js/all.js");

import inputs  from "./inputs.js";
import buybacks  from "./buybacks.js";

$(function() {
    inputs.setUp();
    buybacks.setUp();
})
(function() {
    var injectParams = [];

    var DefaultController = function() {
    };

    DefaultController.$inject = injectParams;

    angular.module('winch').controller('DefaultController', DefaultController);
}());

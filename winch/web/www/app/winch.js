// Definition of the module for the Winch application, with routes 
//    laid out for the views and controllers
//
(function() {
    var app = angular.module('winch', [
        'ngRoute',
        'ngAnimate'
    ]);

    app.config(function($routeProvider, $locationProvider) {
        var viewBase = '/app/views/'

        $routeProvider
        .when('/welcome', {
            templateUrl: viewBase + 'welcome.html',
            controller: 'DefaultController'
        })
        .when('/about', {
            templateUrl: viewBase + 'about.html',
            controller: 'DefaultController'
        })
        .when('/login', {
            templateUrl: viewBase + 'login.html',
            controller: 'LoginController'
        })
        .when('/register', {
            templateUrl: viewBase + 'register.html',
            controller: 'RegistrationController'
        })
        .when('/board/:boardId', {
            templateUrl: viewBase + 'board.html',
            controller: 'BoardController'
        })
        .otherwise({redirectTo: '/welcome'});

        $locationProvider.hashPrefix('winch');
    });
}());

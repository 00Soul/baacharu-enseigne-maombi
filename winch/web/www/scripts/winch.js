// Definition of the module for the Winch application, with routes 
//    laid out for the views and controllers
//
angular.module('winch', ['ngRoute'])
    .controller('DefaultController', function($scope) {
    })
    .controller('LoginController', function($scope) {
    })
    .controller('RegistrationController', function($scope) {
    })
    .controller('BoardController', function($scope) {
    })
    .config(function($routeProvider, $locationProvider) {
        var viewBase = '/views/'

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
            });

        $locationProvider.hashPrefix('winch');
    });

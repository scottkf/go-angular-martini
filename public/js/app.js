'use strict';

var issueTracker = angular.module('issueTracker', ['issues', 'ngRoute']).
config(function($routeProvider, $locationProvider) {
  $routeProvider.when('/', {
    templateUrl: '/js/templates/issues/index.html',
    controller: 'IssueListCtrl'
  })
  $routeProvider.when('/issues/:id', {
    templateUrl: '/js/templates/issues/show.html',
    controller: 'IssueShowCtrl'
  })
  $locationProvider.html5Mode(true);
})
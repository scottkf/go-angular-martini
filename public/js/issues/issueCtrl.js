'use strict';

issues.controller('IssueListCtrl', ['$scope', '$resource', 'Issue', function($scope, $resource, Issue) {
  $scope.issues = Issue.query();
  $scope.issue = new Issue({});

  $scope.cancel = function() {
    $scope.issue = new Issue({});
    $scope.form.$setPristine();
  }

  $scope.create = function(issue) {
    issue.$save(function(i, responseHeaders) {
      $scope.issues.unshift(issue);
      $scope.issue = new Issue({});
    },
    function(result) {
      $scope.form.title.$setValidity("unique", false);
      console.log(result);
      console.log(issue);
    });
  }
}]);

issues.controller('IssueShowCtrl', ['$scope', '$resource', '$routeParams','$location','Issue', function($scope, $resource, $routeParams, $location, Issue) {
  $scope.issue = Issue.get({Id: $routeParams.id});
  $scope.original = {};
  $scope.isEditing = false;

  $scope.edit = function() {
    $scope.isEditing = true;
    $scope.original = angular.copy($scope.issue);
  }

  $scope.cancel = function() {
    $scope.isEditing = false;
    $scope.issue = angular.copy($scope.original);
  }

  $scope.update = function(issue) {
    issue.$update();
    $scope.isEditing = false;
  }

  $scope.destroy = function(issue) {
    Issue.remove({Id: issue.id});
    $location.path('/');
  }
}]);
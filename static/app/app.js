var app = angular.module('app', ['ngRoute'])
  .config(function ($routeProvider, $locationProvider) {
    $routeProvider
      .when('/search', {
        templateUrl: '/static/search.html',
        controller: 'SearchController'
      })
      .otherwise('/');
  })
  .controller('MainController', function ($scope, $location) {
    $scope.searchQuery = "new girl";
    $scope.searchShow = function (q) {
      $location.path("/search").search({
        name: q
      });
    };
  })
  .controller('SearchController', function ($scope, $location, $http) {
    $scope.name = $location.search().name;
    if (!$scope.name) {
      return;
    }
    console.log("searching:", $scope.name);
    $scope.searching = true;
    $http.get("/api/search", {
      params: {
        q: $scope.name
      }
    }).success(function (shows) {
      console.log("found shows:", shows);
      angular.forEach(shows, function (v, k) {
        shows[k].firstAired = moment(v.FirstAired);
        shows[k].actors = v.Actors.splice(0, 5);
      });
      $scope.shows = shows;
    }).error(function (resp) {
      console.log("error", resp);
    }).finally(function () {
      $scope.searching = false;
    });
  });

var fileManagerApp = angular.module('fileManagerApp', []);

fileManagerApp.controller('fileManagerController', function($scope, $http) {

    $scope.listFiles = function(){

        if($scope.path != ""){
            $http({
                url:   "http://localhost:8080/api/getlist",
                method: "GET",
                data:   { path: $scope.path }
            }).then(function(response){
                console.log(response);
            }).catch(function(error){
                console.log(error);
            });
        }
        else{
            alert("Please enter a directory path");
        }
        
    }
    
});
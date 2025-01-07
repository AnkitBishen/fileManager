var fileManagerApp = angular.module('fileManagerApp', []);

fileManagerApp.controller('fileManagerController', function($scope, $http) {

    $scope.listFiles = function(){

        if($scope.path != ""){
            $http({
                url:   "http://localhost:8080/api/getlist",
                method: "POST",
                data:   { path: $scope.path }
            }).then(function(response){
                console.log(response);
                $scope.dirList = response.data.data;
            }).catch(function(error){
                console.log(error);
            });
        }
        else{
            alert("Please enter a directory path");
        }
        
    }
    
});
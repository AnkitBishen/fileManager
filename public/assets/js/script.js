var fileManagerApp = angular.module('fileManagerApp', []);

fileManagerApp.controller('fileManagerController', function($scope, $http) {

    $scope.manage_createDir = {
        name: "",
        type: "folder",
        extension: ".txt"
    }
    $scope.createDir = function(){
        if($scope.manage_createDir.name != ""){
            $http({
                url:   "http://localhost:8080/api/gddetlist",
                method: "POST",
                data:   { data: $scope.manage_createDir },
                header: { 'Content-Type': 'application/json' }
            }).then(function(response){
                console.log(response);
                
                // Close the modal
                const modal = bootstrap.Modal.getInstance(document.getElementById('createModal'));
                modal.hide();
            }).catch(function(error){
                console.log(error);
            });
        }
        else{
            alert("Please enter " + $scope.manage_createDir.type + " name");
        }
    }



    $scope.dirList = [];
    $scope.listFiles = function(){
        if($scope.path != ""){
            $http({
                url:   "http://localhost:8080/api/getlist",
                method: "POST",
                data:   { path: $scope.path },
                header: { 'Content-Type': 'application/json' }
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
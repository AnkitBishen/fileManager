var fileManagerApp = angular.module('fileManagerApp', []);

fileManagerApp.controller('fileManagerController', function($scope, $http) {

    $scope.manage_createDir = {
        name: "",
        type: "folder",
        extension: ".txt",
        currentDirPath: ""
    }
    $scope.createDir = function(){
        if($scope.manage_createDir.name != ""){
            $http({
                url:   "http://localhost:8080/api/createDir",
                method: "POST",
                data:   $scope.manage_createDir ,
                header: { 'Content-Type': 'application/json' }
            }).then(function(response){
                console.log(response);

                if(response.data.status == "success"){
                    alert(response.data.data)
                    $scope.listFiles();
                }
                else{
                    alert(response.data.data)
                }
                
                // Close the modal
                const modal = bootstrap.Modal.getInstance(document.getElementById('createModal'));
                modal.hide();
            }).catch(function(error){
                if(error.status == 403){    
                    alert(error.data.error)
                }
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
                $scope.manage_createDir.currentDirPath = $scope.path
            }).catch(function(error){
                console.log(error);
            });
        }
        else{
            alert("Please enter a directory path");
        }
        
    }
    
});
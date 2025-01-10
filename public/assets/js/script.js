var fileManagerApp = angular.module('fileManagerApp', []);

fileManagerApp.controller('fileManagerController', function($scope, $http) {

    // create file and folders
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
                // console.log(response);

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


    // list sub folders and files
    $scope.dirList = [];
    $scope.listFiles = function(){
        if($scope.path != ""){
            $http({
                url:   "http://localhost:8080/api/getlist",
                method: "POST",
                data:   { path: $scope.path },
                header: { 'Content-Type': 'application/json' }
            }).then(function(response){
                // console.log(response);
                $scope.dirList = response.data.data;
                $scope.manage_createDir.currentDirPath = $scope.path

                if($scope.perviousFolderPaths.length == 0){
                    $scope.perviousFolderPaths.push($scope.path)
                }

            }).catch(function(error){
                // console.log(error);
            });
        }
        else{
            alert("Please enter a directory path");
        }
        
    }
   
    
    // navigate to sub folder and pervious folder
    $scope.backBtn = false;
    $scope.perviousFolderPaths = [];
    $scope.moveCount = 0
    $scope.moveToFolder = function(action, path){

        if(action == 'in'){
            $scope.backBtn = true;
            $scope.perviousFolderPaths.push(path)
            $scope.path = path
            $scope.moveCount++

        }else{
            if($scope.perviousFolderPaths.length > 1){
                $scope.perviousFolderPaths.pop()
            }
            $scope.path = $scope.perviousFolderPaths.slice(-1)[0]
            $scope.moveCount--
        }

        if($scope.moveCount == 0){
            $scope.backBtn = false;
        }

        
        // load sub folders
        $scope.listFiles()
    }


    // delete file or folder
    $scope.deleteDir = function(path){
        $http({
            url:   "http://localhost:8080/api/delete",
            method: "POST",
            data:   { currentDirPath: path },
            header: { 'Content-Type': 'application/json' }
        }).then(function(response){
            if(response.data.status == "success"){
                alert(response.data.data)
                $scope.listFiles();
            }
            else{
                alert(response.data.data)
            }
        }).catch(function(error){
            // console.log(error);
        });
    }


    // get properties
    $scope.getProperties = function(dir){
        $scope.propertyData = dir
        $scope.new_dirName = dir.name
    }

    $scope.renameDir = function(dir){

         // rename file or folder
         if($scope.new_dirName.trim() != dir.name.trim()){
            // create new path of new name
            var newPath = dir.path.replace(dir.name, $scope.new_dirName)

            var details = {
                oldName: dir.path,
                newName: newPath
            }

            $http({
                url:   "http://localhost:8080/api/rename",
                method: "POST",
                data:   details,
                header: { 'Content-Type': 'application/json' }
            }).then(function(response){
                if(response.data.status == "success"){
                    alert(response.data.data)
                    $scope.listFiles();


                    // Close the modal
                    const propModal = bootstrap.Modal.getInstance(document.getElementById('propertiesModal'));
                    propModal.hide();
                }
                else{
                    alert(response.data.data)
                }
            }).catch(function(error){
                // console.log(error);
            });
        }

    }
});
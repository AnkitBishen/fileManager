<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
   
</head>
<body>
    <h1>File Manager API</h1>
    <p>A powerful and simple file manager application built with Golang, offering RESTful APIs to manage files and directories. This application allows users to navigate directories, create, rename, delete files and folders, and fetch detailed metadata.</p>
    <h2>Features</h2>
    <ul>
        <li>List Files and Folders: Retrieve a list of all files and folders in a specified directory.</li>
        <li>Create Files and Folders: Easily create new files or directories within a specified path.</li>
        <li>Rename Files and Folders: Rename files or directories as needed.</li>
        <li>Delete Files and Folders: Remove files or directories from the specified path.</li>
        <li>View File Details:
            <ul>
                <li>File size</li>
                <li>File path</li>
                <li>Permissions</li>
                <li>File extension</li>
                <li>Date modified</li>
            </ul>
        </li>
        <li>Navigate Subfolders: Seamlessly move into subfolders and explore their contents.</li>
    </ul>
    <h2>Tech Stack</h2>
    <ul>
        <li><strong>Language:</strong> Go</li>
        <li><strong>API Framework:</strong> Built-in HTTP package</li>
        <li><strong>Architecture:</strong> RESTful APIs</li>
    </ul>
    <h2>Setup and Installation</h2>
    <h3>Prerequisites</h3>
    <ul>
        <li>Go (version 1.18 or later) installed on your machine.</li>
        <li>A terminal or command-line tool.</li>
    </ul>
    <h3>Steps to Run</h3>
    <ol>
        <li>Clone the repository:
            <pre><code>git clone https://github.com/AnkitBishen/fileManager.git 
cd fileManager</code></pre>
        </li>
        <li>Build and run the application:
            <pre><code>go run ./cmd/main.go</code></pre>
        </li>
        <li>Access the API using tools like Postman or <code>curl</code>.</li>
    </ol>
    <h2>API Endpoints</h2>
    <h3>1. List Files and Folders</h3>
    <p><strong>Endpoint:</strong> <code>POST /api/getlist</code></p>
    <div >
        <strong>Request Body:</strong>
        <pre><code>{
  "path": "your-directory"
}</code></pre>
    </div>
    <h3>2. Create File/Folder</h3>
    <p><strong>Endpoint:</strong> <code>POST /api/createDir</code></p>
    <p><strong>Request Body:</strong></p>
    <pre><code>{
        "name": "mFolder 04",
        "type": "file",
        "extension": ".txt",
        "currentDirPath": "your-directory"
    }</code></pre>
    <h3>3. Rename File/Folder</h3>
    <p><strong>Endpoint:</strong> <code>POST /api/rename</code></p>
    <p><strong>Request Body:</strong></p>
    <pre><code>{
    "oldName": "your-directory-path\\mFolder 042",
    "newName": "your-directory-path\\mFolder 042-new"
}</code></pre>
    <h3>4. Delete File/Folder</h3>
    <p><strong>Endpoint:</strong> <code>POST /api/delete</code></p>
    <p><strong>Request Body:</strong></p>
    <pre><code>{
    "currentDirPath": "your-directory-path\\mFolder 042"
}</code></pre>
    <h2>Future Enhancements</h2>
    <ul>
        <li>Implement user authentication.</li>
        <li>Add support for file uploads and downloads.</li>
        <li>Include pagination for large directories.</li>
        <li>Extend support for file previews (text/image files).</li>
    </ul>
    <h2>License</h2>
    <p>This project is licensed under the MIT License. See the <a href="LICENSE">LICENSE</a> file for more details.</p>
</body>
</html>

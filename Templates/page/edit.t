<html>
    <meta>
                    <meta charset="UTF-8">
                    <script src="https://static.imovlr.com/static/jquery-2.2.2.min.js"></script>
                    <script src="https://static.imovlr.com/static/jquery.form.js"></script>
                    <script src="https://static.imovlr.com/static/lib/codemirror.js"></script>
                    <link rel="stylesheet" href="https://static.imovlr.com/static/lib/codemirror.css">
                    <script src="https://static.imovlr.com/static/mode/javascript/javascript.js"></script>
                    <script src="https://static.imovlr.com/static/editor_timeout.js"></script>
                    <link rel="stylesheet" href="https://static.imovlr.com/static/full_editor.css" type="text/css">
    </meta>
<body>
    <form id="editor_form" method="POST" action="/edit/{{ .Filename }}">
        <textarea id="editor" name="content">{{ printf "%s" .Body }}</textarea>
    </form>
    
    <script>
        function processResponse ( response ) {
            console.log( 'Saved correctly {{ .Filename }}' );
            document.body.style.background = '#BFB';
            savedTimeout = setTimeout( function(){
                document.body.style.background = '#FFF';
            } , 300  );
        }
        function preProcess ( formData ){
            // Nothing to pre process
        }
        function errorProcess ( formData ) {
            console.log( 'An Error happened when saving the file' );
            document.body.style.background = '#FBB';
        }
        var editorTimeout;
        $('#editor_form').ajaxForm({
            beforeSubmit: preProcess,
            success: processResponse,
            error: errorProcess,
        });
        var contentcode    = CodeMirror.fromTextArea(
                                document.getElementById("editor"),
                                {
                                    mode: {
                                            name: "text/x-perl",
                                            globalVars: true
                                    },
                                    lineNumbers: true,
                                    lineWrapping: true,
                                    matchBrackets: true,
                                    indentUnit: 4,
                                    indentWithTabs: true
                                }
                            );
        contentcode.on( "change", function() {
            console.log('Triggered on:' + editorTimeout);
            if ( editorTimeout != undefined ) clearTimeout( editorTimeout );
            editorTimeout = setTimeout( function(){ document.body.style.background = '#FFB'; } , 1000  );
            editorTimeout = setTimeout( function(){
                    document.getElementById("editor").value = contentcode.doc.getValue();
                    document.body.style.background = '#BBB';
                    $('#editor_form').submit();
                },
                3000  );
        });
    </script>

</body>

</html>

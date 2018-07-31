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
    <form id="editor_form" method="POST" action={{ .Filename }}>
        <textarea id="editor" name="content">{{ printf "%s" .Body }}</textarea>
        <input type="submit">
        <a href="/">[ List of Files ]</a>
        <!-- Plugins --!>
    </form>
    
    <script>
        function processResponse ( response ) {
            console.log( 'hiro response' + response );
            var template = document.createElement('template');
            template.innerHTML = response;
            var status =  template.content.getElementById("boom").innerHTML;
            if ( status == 1 ) {
                document.body.style.background = '#BFB';
            }else{
                document.body.style.background = '#FBB';
            }
        }
        function preProcess ( formData){
            //console.log( 'Form:' );
            //console.log(  JSON.stringify(formData ) );
            //console.log( 'Here goes:');
            //console.log( contentcode.doc.getValue() );
            //formData = contentcode.doc.getValue();
        }
        var editorTimeout;
        $('#editor_form').ajaxForm({
            beforeSubmit: preProcess,
            success: processResponse,
            error: console.log( 'Not today' )
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
            console.log('hiro on' + editorTimeout);
            if ( editorTimeout != undefined ) clearTimeout( editorTimeout );
            editorTimeout = setTimeout( function(){ document.body.style.background = '#FFB'; } , 1000  );
            editorTimeout = setTimeout( function(){
                    document.getElementById("editor").value = contentcode.doc.getValue();
                    document.body.style.background = '#BBB';
                    //console.log('before');
                    $('#editor_form').submit();
                    //console.log('afetr');
                    //document.body.style.background = '#BFB';
                },
                3000  );
    
    //        editorTimeout = setTimeout( function(){ alert( 'Hiro de Janeiro'); } , 15000  );
    
    
        });
    </script>

</body>

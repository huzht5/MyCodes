$(document).ready(function() {
    $.ajax({
        url: "/Function2/test"
    }).then(function(data) {
       $('.greeting-id').append(data.id);
       $('.greeting-content').append(data.content);
    });
});

$(document).on('ready', function() {
  $.getJSON('/api/issues', function(data) {
    if (data == null) {
      return
    }
    $.each(data, function(index, value) {
      $('#issues').append(createIssue(value))
    })
  })

  var o = {
    name: "Hello"
  }

  var f = function() {
    console.log(this);
  }
  f()
  f.apply(o, [1,2,3])
  f.call(o, 1, 2, 3)
  f.bind(this)()


  $(document).on('click', '.delete', function(e) {
    $this = $(this)
    $.ajax({
      url: '/api/issues/' + $(this).parent().data('id'),
      type: 'DELETE',
      complete: function(data) {
        if (data.status == 204) {
          $this.parent().remove();
        } else {
          alert('Not removed!');
        }
      }
    })
  })


  $(document).on('submit', 'form', function(e) {
    e.preventDefault();
    $searchValue = $(this).find('input[type=text]').val()
    // Validation
    if ($searchValue == "") {
      return
    }
    $.getJSON('/api/issues?title='+$searchValue, function(data) {
      $('#issues').html('');
      if (data == null) {
        return
      }
      $.each(data, function(index, value) {
        $('#issues').append(createIssue(value))
      });
    })
  })
});


function createIssue(issue) {
  return $('<div class="issue" data-id="'+issue.id+'">'+issue.title+': '+issue.body+'<button class="delete">remove</button></div>')
}
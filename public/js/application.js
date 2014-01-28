$(document).on('ready', function() {
  $.getJSON('/issues', function(data) {
    if (data == null) {
      return
    }
    $.each(data, function(index, value) {
      $('#issues').append(createIssue(value))
    })
  })

  $(document).on('click', '.delete', function(e) {
    $this = $(this)
    $.ajax({
      url: '/issues/' + $(this).parent().data('id'),
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
    $.getJSON('/issues?title='+$searchValue, function(data) {
      $('#issues').html('');
      if (data == null) {
        return
      }
      $.each(data, function(index, value) {
        $('#issues').append(createIssue(value))
      });
    })
  })
})


function createIssue(issue) {
  return $('<div class="issue" data-id="'+issue.id+'">'+issue.title+': '+issue.body+'<button class="delete">remove</button></div>')
}
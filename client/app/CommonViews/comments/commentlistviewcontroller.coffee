class CommentListViewController extends KDListViewController
  constructor:->
    super
    @startListeners()

  instantiateListItems:(items)->

    newItems = []

    for comment, i in items
      
      nextComment = items[i+1]

      skipComment = no
      if nextComment? and comment.deletedAt
        if Date.parse(nextComment.meta.createdAt) > Date.parse(comment.deletedAt)
          skipComment = yes
      
      if not nextComment and comment.deletedAt
        skipComment = yes

      unless skipComment
        commentView = @getListView().addItem comment
        newItems.push commentView
    
    return newItems

  startListeners:->
    listView = @getListView()
    
    listView.on 'ItemWasAdded', (view, index)=>
      view.on 'CommentIsDeleted', ->
        view.setClass "deleted"
        view.$().html "<div class='item-content-comment clearfix'><span>This comment has been deleted.</span></div>"
    
    listView.on "AllCommentsLinkWasClicked", (commentHeader)=>

      # some problems when logged out server doesnt responds
      @utils.wait 5000, -> listView.emit "BackgroundActivityFinished"
      
      @fetchAllComments 0, (err, comments)=>
        @removeAllItems()
        @instantiateListItems comments

    listView.registerListener
      KDEventTypes  : "CommentSubmitted"
      listener      : @
      callback      : (pubInst, reply)->
        model = listView.getData()
        listView.emit "BackgroundActivityStarted"
        model.reply reply, (err, reply)->
          # listView.emit "AllCommentsLinkWasClicked"
          listView.addItem reply
          listView.emit "OwnCommentHasArrived"
          listView.emit "BackgroundActivityFinished"

  fetchCommentsByRange:(from,to,callback)=>
    [to,callback] = [callback,to] unless callback
    query = {from,to}
    message = @getListView().getData()
    message.commentsByRange query,(err,comments)=>
      @getListView().emit "BackgroundActivityFinished"
      callback err,comments
  
  fetchAllComments:(skipCount=3, callback = noop)=>
    
    listView = @getListView()
    listView.emit "BackgroundActivityStarted"
    message = @getListView().getData()
    message.restComments skipCount, (err, comments)=>
      listView.emit "BackgroundActivityFinished"
      listView.emit "AllCommentsWereAdded"
      callback err, comments
  
  replaceAllComments:(comments)->
    @removeAllItems()
    @instantiateListItems comments

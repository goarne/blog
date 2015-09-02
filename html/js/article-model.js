/*
The article js 
*/
var articleURL = "http://localhost:8888/blog/article/";


var Article = Backbone.Model.extend({
	defaults:{
		ID: "",
		Title : "",
		Abstract : "",
		Text: "",
		Date: "",
		Author:""},
	idAttribute:"ID",
	urlRoot:  articleURL
});

var ArticleCollection = Backbone.Collection.extend({
	model:Article,
	url:  articleURL,
	initialize: function(){

		/*When fetching your collection, the reset event is not fired by default anymore. 
		In order to have Backbone fire the reset event when the collection has been fetched as done below.
		articles.fetch({reset: true});*/
		// Assign the Deferred issued by fetch() as a property
		this.deferred = this.fetch({reset: true}); 
	}
});



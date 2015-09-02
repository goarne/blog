var ArticleAbstractView = Backbone.View.extend({
	model: Article,
	tagname: "div",               

	events: {
		"click #show-article": "showArticle"
	},

	showArticle:function(e){
		e.preventDefault();
		//Element clicked on

		var target =  $(e.currentTarget);
		window.location.href= "/post.html?id=" + this.model.get("Id");
	},

	render: function (){						
		this.$el.html('<div class="post-preview">' +
						'<a href="#" id="show-article">' +
							'<h2 class="post-title">' +
								this.model.get("Title") +
							'</h2>' +
							'<h3 class="post-subtitle">' +
								this.model.get("Abstract") +
							'</h3>' +
						'</a>' +
						'<p class="post-meta">Posted by <a href="#">' + this.model.get("Author") +'</a> on ' + this.model.get("Date") + '</p>' +
					'</div>');

		return this;
	}
});

var ArticleListView = Backbone.View.extend({
	model: ArticleCollection,

	render: function() {
		this.$el.html(); // lets render this view        
		var self = this;

		this.model.deferred.done(function() {

			for(var i = 0; i < self.model.length; i++) {
				// lets create a book view to render
				var article = self.model.at(i)

				if(article != null){
					var artView = new ArticleAbstractView({model: article});
					self.$el.append(artView.$el); 
					artView.render(); 
				}
			} 

			return self;
		}); 
	},
});

var ArticleHeaderView = Backbone.View.extend({
	model: Article,

	render: function (){
		this.$el.html(
			'<h1>' + this.model.get("Title") + '</h1>' +
			'<h2 class="subheading">'+ this.model.get("Abstract") +'</h2>' +
			'<span class="meta">Posted by ' + this.model.get("Author") + '</span>');
		return this;
	}
});

var ArticleBodyView = Backbone.View.extend({
	model: Article,

	render: function (){
		this.$el.html(this.model.get("Text"));
		return this;
	}
});

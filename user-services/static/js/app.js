ReactDOM.render(
    React.createElement(GraphiQL, {
      fetcher: GraphiQL.createFetcher({
        url: window.location.origin + '/graphql',
      }),
      defaultEditorToolsVisibility: true,
      query:
        `#Welcome to the GraphiQL IDE for the Course Services API.`  
    }),
    document.getElementById('graphiql'),
  );
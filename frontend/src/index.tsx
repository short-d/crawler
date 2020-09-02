import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './component/App';
import * as serviceWorker from './serviceWorker';
import {GraphQLService} from './service/fw/GraphQL.service';
import {FetchHTTPService} from './service/fw/http.service';
import {SampleAppGraphqlService} from './service/sampleapp.graphql.service';
import {EnvService} from './service/env.service';

const envService = new EnvService();
const fetchHTTPService = new FetchHTTPService();
const graphQLService = new GraphQLService(fetchHTTPService);
const sampleAppGraphQLService = new SampleAppGraphqlService(graphQLService, envService);

ReactDOM.render(
  <React.StrictMode>
    <App
      envService={envService}
      sampleAppGraphQLService={sampleAppGraphQLService}
    />
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();

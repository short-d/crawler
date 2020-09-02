import {GraphQLService, IGraphQLQuery} from './fw/GraphQL.service';
import {EnvService} from './env.service';
import {ChangeLog} from '../entity/changelog';

interface IGraphQLQueryResult {
  authQuery: IGraphQLAuthQueryResult;
}

interface IGraphQLAuthQueryResult {
  changeLog: IGraphQLChangeLog;
}

interface IGraphQLChangeLog {
  changes: IGraphQLChange[];
}

interface IGraphQLChange {
  id: string;
  title: string;
}

export class SampleAppGraphqlService {
  private readonly baseURL: string;

  constructor(private graphQLService: GraphQLService, private envService: EnvService) {
    this.baseURL = `${envService.getEnv().graphQLBaseURL}/graphql`;
  }

  getChangeLog(): Promise<ChangeLog> {
    const query: IGraphQLQuery = {
      query: `
query {
  authQuery {
    changeLog {
      changes {
        id
        title
      }
    }
  }
}
      `,
      variables: {
        authToken: ''
      }
    };
    return this
      .graphQLService
      .query<IGraphQLQueryResult>(this.baseURL, query)
      .then((result: IGraphQLQueryResult) => result.authQuery.changeLog);
  }
}
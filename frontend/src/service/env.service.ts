export interface Environment {
  env: string;
  graphQLBaseURL: string;
}

export class EnvService {
  private readonly environment: Environment;

  constructor() {
    this.environment = {
      env: this.getVal('ENV'),
      graphQLBaseURL: this.getVal('GRAPHQL_API_BASE_URL')
    }
  }

  getEnv(): Environment {
    return this.environment;
  }

  private getVal(name: string): string {
    return process.env[`REACT_APP_${name}`] || '';
  }
}
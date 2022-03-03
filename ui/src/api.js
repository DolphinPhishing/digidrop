import { ApolloClient, InMemoryCache } from '@apollo/client';
import { createUploadLink } from 'apollo-upload-client';

const client = new ApolloClient({
  cache: new InMemoryCache(),
  link: createUploadLink({ uri: `${process.env.REACT_APP_API_URI}/query`, credentials: 'include' }),
  credentials: 'include',
});

export default client;

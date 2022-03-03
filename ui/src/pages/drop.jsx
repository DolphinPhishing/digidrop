import React, { useState, useRef, useEffect } from 'react';
import { useCookies } from 'react-cookie';
import { useNavigate } from 'react-router-dom';
import { gql, useMutation, useApolloClient } from '@apollo/client';

import './drop.scss';

const UPLOAD_MUTATION = gql`
  mutation upload($file: Upload!) {
    upload(file: $file) {
      id
      url_id
    }
  }
`;

const Drop = () => {
  const [isAuthed, setIsAuthed] = useState(false);
  const [authLoading, setAuthLoading] = useState(false);
  const [uploadFileMutation] = useMutation(UPLOAD_MUTATION);
  const apolloClient = useApolloClient();
  const [token, setToken] = useState('');
  const navigate = useNavigate();
  const [response, setResponse] = useState(null);

  useEffect(() => {
    fetch(`${process.env.REACT_APP_API_URI}/me`, {
      credentials: 'include',
    }).then((res) => {
      setIsAuthed(res.status === 200);
    }, console.error);
  }, [authLoading]);

  const onChange = ({
    target: {
      validity,
      files: [file],
    },
  }) => validity.valid
    && uploadFileMutation({ variables: { file } }).then((res) => {
      apolloClient.resetStore();
      if (res.data && res.data.upload) {
        setResponse(res.data.upload);
      }
    }, console.error);

  const storeToken = () => {
    const formData = new FormData();
    formData.append('token', token);
    setAuthLoading(true);
    fetch(`${process.env.REACT_APP_API_URI}/auth`, {
      method: 'post',
      credentials: 'include',
      body: formData,
    }).then((res) => {
      if (res.status !== 200) console.error('Unknown error occurred. Please contact T H E  K E E P E R');
      setAuthLoading(false);
    }, console.error);
  };

  const clearUpload = () => {
    setResponse(null);
  };

  return (
    <div className="drop-container">
      {!isAuthed ? (
        <div className="token-container">
          <input
            type="text"
            className="token"
            placeholder="Your token please..."
            autoFocus
            onChange={(e) => {
              e.preventDefault();
              setToken(e.target.value);
            }}
          />
          <button type="button" className="token-submit" onClick={storeToken} disabled={!token}>&gt;</button>
        </div>
      )
        : (
          <div>
            <h1>🕵️ DigiDrop 🕵️</h1>
            <p><em>For your eyes only</em></p>
            {response ? (
              <div className="confirmation-code">
                <p>Your confirmation code:</p>
                <code className="code">{response.url_id}</code>
                <p><small><em>Please call in your drop to 88888</em></small></p>
                <button type="button" className="upload-another" onClick={clearUpload}>Upload another</button>
              </div>
            ) : (
              <input type="file" className="file-upload" multiple required onChange={onChange} />
            )}
          </div>
        )}
    </div>
  );
};

export default Drop;

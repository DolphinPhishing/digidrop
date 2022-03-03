import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useCookies } from 'react-cookie';
import { gql, useQuery } from '@apollo/client';

import './admin.scss';

const LIST_FILES_QUERY = gql`
 query ListFiles {
    files {
      id
      url_id
      file_path
      accessed
      user
    }
  }
`;

const Admin = () => {
  const [cookies, setCookie, removeCookie] = useCookies(['dead-auth']);
  const [user, setUser] = useState(null);
  const navigate = useNavigate();
  const {
    data: listFilesData,
    error: listFilesError,
    loading: listFilesLoading,
  } = useQuery(LIST_FILES_QUERY);
  const [filteredFileList, setFilteredFileList] = useState([]);
  const [searchValue, setSearchValue] = useState('');

  useEffect(() => {
    if (!searchValue && listFilesData) setFilteredFileList([...listFilesData.files]);
  }, [listFilesData]);

  useEffect(() => {
    fetch(`${process.env.REACT_APP_API_URI}/me`, {
      credentials: 'include',
    }).then((res) => res.json()).then((u) => {
      if (u.type !== 'ADMIN') {
        navigate('/drop');
      }
      setUser(u);
    }).catch(console.error);
  }, [cookies['dead-auth']]);

  const onSearchChange = (e) => {
    if (!listFilesData.files) return;
    setSearchValue(e.target.value);
    if (!e.target.value) setFilteredFileList(listFilesData.files);
    else {
      setFilteredFileList(
        [...listFilesData.files]
          .filter((f) => f.url_id.toLowerCase().includes(e.target.value.toLowerCase())),
      );
    }
  };

  return (
    <div className="admin-container">
      <div className="header">
        <h2>DigiDrop Admin Panel</h2>
        <p>
          Welcome,&nbsp;
          {user && user.name}
        </p>
      </div>
      <div className="file-search">
        <input type="text" placeholder="Search for confirmation number..." value={searchValue} onChange={onSearchChange} />
      </div>
      <div className="file-list">
        <table>
          <thead>
            <tr>
              <th>Confirmation #</th>
              <th>User</th>
              <th>Accessed?</th>
              <th />
            </tr>
          </thead>
          <tbody>
            {filteredFileList && filteredFileList.map((file) => (
              <tr key={file.id}>
                <td>{file.url_id}</td>
                <td>{file.user}</td>
                <td>{file.accessed ? 'true' : 'false'}</td>
                <td><a href={`${process.env.REACT_APP_API_URI}/download/${file.url_id}`}>Download</a></td>
              </tr>
            ))}
            {filteredFileList.length === 0 && (
              <tr>
                <td style={{ textAlign: 'center' }} colSpan={4}>No files uploaded</td>
              </tr>
            )}
          </tbody>
        </table>
      </div>
    </div>
  );
};
export default Admin;

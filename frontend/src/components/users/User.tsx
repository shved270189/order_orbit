import { useState, useEffect } from 'react'
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Moment from 'moment';
import {useParams} from "react-router";
import axios from 'axios';
import BackButton from '../common/BackButton';

type User = {
  id: number;
  fullName: string;
  login: string;
  updatedAt: string;
  createdAt: string;
};

function User() {
  const [user, setUser] = useState<User | null>(null)
  const { userId } = useParams<{ userId: string }>()


  useEffect(() => {
    const fetchUser = () => {
      axios.get(`${import.meta.env.VITE_API_BASE_URL}/users/${userId}`)
        .then(response => response.data)
        .then(data => setUser(data))
        .catch(error => console.error('Error fetching data:', error))
    }
    fetchUser()
  }, [userId])

  return (
    <>
    <BackButton to="/users" />
    <Row>
      <Col>
      {user === null && (
        <div>
          <h2>Loading...</h2>
        </div>
      )}
      {user && (
        <div>
          <h2>{user.fullName}</h2>
          <p>Login: {user.login}</p>
          <p>Updated At: {Moment(user.updatedAt).format('MMMM Do YYYY, h:mm:ss a')}</p>
          <p>Created At: {Moment(user.createdAt).format('MMMM Do YYYY, h:mm:ss a')}</p>
        </div>
      )}
      </Col>
    </Row>
    </>
  )
}

export default User

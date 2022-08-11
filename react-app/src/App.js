import { useEffect, useState } from 'react';
import './App.css';
import axios from 'axios';

function App() {

  const [notification, setNotification] = useState('');

  const handlePing = async () => {
      try {
          const response = await axios.get('/api/teste');
          setNotification(`Successful ping with response: ${response.data}`);
      } catch (e) {
          setNotification('Failed to ping');
      }
  }

  useEffect(() => {
    handlePing()
  }, [])

  return (
      <div>
          <div>
              <p>Notification: {notification}</p>
          </div>
      </div>
  );
}

export default App;

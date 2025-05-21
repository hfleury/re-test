import React, { useState } from 'react';

const API_BASE_URL = process.env.REACT_APP_API_BASE_URL || 'http://localhost:8081';

function App() {
  const [orderAmount, setOrderAmount] = useState('');
  const [result, setResult] = useState(null);
  const [error, setError] = useState(null);

  const handleCalculate = async () => {
    setError(null);
    setResult(null);

    if (!orderAmount || isNaN(orderAmount) || Number(orderAmount) <= 0) {
      setError('Please enter a valid positive integer for order amount.');
      return;
    }

    try {
      const response = await fetch(`${API_BASE_URL}/calculate?order_amount=${orderAmount}`);
      const data = await response.json();

      if (!response.ok) {
        setError(data.message || 'Something went wrong');
        return;
      }

      setResult(data.data.packs);
    } catch (err) {
      setError('Failed to fetch data from the API.');
    }
  };

  return (
    <div style={{ margin: '2rem' }}>
      <h1>Pack Size Calculator</h1>
      <input
        type="number"
        placeholder="Enter order amount"
        value={orderAmount}
        onChange={(e) => setOrderAmount(e.target.value)}
      />
      <button onClick={handleCalculate} style={{ marginLeft: '1rem' }}>
        Calculate
      </button>

      {error && <p style={{ color: 'red' }}>{error}</p>}

      {result && (
        <div>
          <h2>Result Packs:</h2>
          <ul>
            {Object.entries(result).map(([packSize, count]) => (
              <li key={packSize}>
                Pack size {packSize}: {count}
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
}

export default App;

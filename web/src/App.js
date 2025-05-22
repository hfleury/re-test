import React, { useState } from 'react';

const API_BASE_URL = process.env.REACT_APP_API_BASE_URL || 'http://localhost:8081';

function App() {
  console.log('API_BASE_URL:', API_BASE_URL);

  const [orderAmount, setOrderAmount] = useState('');
  const [result, setResult] = useState(null);
  const [error, setError] = useState(null);
  const [packSizes, setPackSizes] = useState(['']);

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

  const handlePackSizeChange = (index, value) => {
    const updated = [...packSizes];
    updated[index] = value;
    setPackSizes(updated);
  };

  const addPackSize = () => {
    setPackSizes([...packSizes, '']);
  };

  const removePackSize = (index) => {
    const updated = packSizes.filter((_, i) => i !== index);
    setPackSizes(updated);
  };

  const submitPackSizes = async () => {
    setError(null);
    try {
      const sizes = packSizes
        .map((size) => parseInt(size, 10))
        .filter((size) => !isNaN(size) && size > 0);

      const response = await fetch(`${API_BASE_URL}/config/pack_sizes`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ pack_sizes: sizes }),
      });

      const data = await response.json();

      if (!response.ok) {
        setError(data.message || 'Failed to update pack sizes');
        return;
      }

      alert('Pack sizes updated successfully!');
    } catch (err) {
      setError('Failed to send pack sizes to the API.');
    }
  };

  return (
    <div style={{ margin: '2rem', maxWidth: '600px' }}>
      <h1>Order Packs Calculator</h1>

      {/* Order Amount Calculator */}
      <div style={{ marginBottom: '1rem' }}>
        <input
          type="number"
          placeholder="Enter order amount"
          value={orderAmount}
          onChange={(e) => setOrderAmount(e.target.value)}
        />
        <button onClick={handleCalculate} style={{ marginLeft: '1rem' }}>
          Calculate
        </button>
      </div>

      {/* Error Message */}
      {error && <p style={{ color: 'red' }}>{error}</p>}

      {/* Calculation Result */}
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

      {/* Pack Size Editor */}
      <div style={{ marginTop: '3rem' }}>
        <h2>Pack Sizes</h2>
        {packSizes.map((size, index) => (
          <div key={index} style={{ marginBottom: '0.5rem' }}>
            <input
              type="number"
              value={size}
              placeholder="Enter pack size"
              onChange={(e) => handlePackSizeChange(index, e.target.value)}
            />
            {packSizes.length > 1 && (
              <button onClick={() => removePackSize(index)} style={{ marginLeft: '0.5rem' }}>
                -
              </button>
            )}
          </div>
        ))}
        <button onClick={addPackSize} style={{ marginTop: '0.5rem' }}>
          + Add Pack Size
        </button>

        <div style={{ marginTop: '1rem' }}>
          <button onClick={submitPackSizes}>Submit Pack Size Changes</button>
        </div>
      </div>
    </div>
  );
}

export default App;

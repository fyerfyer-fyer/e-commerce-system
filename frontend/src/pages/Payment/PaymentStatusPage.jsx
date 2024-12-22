import React from 'react';
import { Card, Typography, Spin } from 'antd';

const { Title, Paragraph } = Typography;

function PaymentStatusPage({ match }) {
  const paymentId = match.params.paymentId; // Assuming you're using react-router v5

  // Simulate fetching payment status from the backend
  const [loading, setLoading] = React.useState(true);
  const [status, setStatus] = React.useState(null);

  React.useEffect(() => {
    // Simulate an API call to fetch payment status
    setTimeout(() => {
      setLoading(false);
      setStatus({
        paymentId: paymentId,
        status: 'completed',
        transactionId: 'TX123456789',
        amount: 69.98,
      });
    }, 1000);
  }, [paymentId]);

  return (
    <div className="page-container">
      <Title level={2}>Payment Status</Title>
      <Card title={`Payment #${paymentId}`} bordered={false}>
        {loading ? (
          <Spin tip="Loading..." size="large" />
        ) : (
          <>
            <Paragraph><strong>Status:</strong> {status.status}</Paragraph>
            <Paragraph><strong>Transaction ID:</strong> {status.transactionId}</Paragraph>
            <Paragraph><strong>Amount:</strong> ${status.amount.toFixed(2)}</Paragraph>
          </>
        )}
      </Card>
    </div>
  );
}

export default PaymentStatusPage;
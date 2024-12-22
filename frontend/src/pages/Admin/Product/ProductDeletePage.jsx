import React from 'react';
import { Typography, Button, Modal, Space, message } from 'antd';
import { useParams, useNavigate, Link } from 'react-router-dom';

const { Title } = Typography;

function ProductDeletePage() {
  const { productId } = useParams();
  const navigate = useNavigate();

  const handleDelete = () => {
    // Here you can add the logic to delete the product on the backend
    console.log('Deleting product:', productId);
    message.success('Product deleted successfully!');
    navigate('/admin/product/list');
  };

  return (
    <div className="page-container">
      <Title level={2}>Delete Product</Title>
      <Modal
        title="Confirm Delete"
        visible={true}
        onOk={handleDelete}
        onCancel={() => navigate('/admin/product/list')}
        okText="Delete"
        okType="danger"
        cancelText="Cancel"
      >
        <p>Are you sure you want to delete the product with ID {productId}?</p>
      </Modal>
    </div>
  );
}

export default ProductDeletePage;
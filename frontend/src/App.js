import React, { useEffect, useState } from "react";
import axios from "axios";
import './App.css'; 
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

function App() {
  const [produtos, setProdutos] = useState([]);
  const [produtoEditando, setProdutoEditando] = useState(null);
  const [mostrarFormulario, setMostrarFormulario] = useState(false);

  const [formulario, setFormulario] = useState({
    nome: '',
    preco: '',
    quantidade: '',
    descricao: ''
  });

  useEffect(() => {
    listarProdutos();
  }, []);

  const listarProdutos = () => {
    axios.get("http://localhost:8080/produtos")
      .then(res => setProdutos(res.data))
      .catch(() => toast.error("❌ Erro ao buscar produtos"));
  };

  const limparFormulario = () => {
    setFormulario({ nome: '', preco: '', quantidade: '', descricao: '' });
    setProdutoEditando(null);
    setMostrarFormulario(false);
  };

  const handleChange = (e) => {
    setFormulario({
      ...formulario,
      [e.target.name]: e.target.value
    });
  };

  const criarProduto = () => {
    const produtoData = {
      ...formulario,
      preco: parseFloat(formulario.preco),
      quantidade: parseInt(formulario.quantidade)
    };

    axios.post("http://localhost:8080/produtos", produtoData)
      .then(() => {
        toast.success("✅ Produto criado com sucesso!");
        listarProdutos();
        limparFormulario();
      })
      .catch(() => toast.error("❌ Erro ao criar produto"));
  };

  const iniciarEdicao = (produto) => {
    setFormulario({
      nome: produto.nome,
      preco: produto.preco.toString(),
      quantidade: produto.quantidade.toString(),
      descricao: produto.descricao
    });
    setProdutoEditando(produto);
    setMostrarFormulario(true);
  };

  const atualizarProduto = () => {
    const produtoData = {
      ...formulario,
      preco: parseFloat(formulario.preco),
      quantidade: parseInt(formulario.quantidade)
    };

    axios.put(`http://localhost:8080/produtos/${produtoEditando.id}`, produtoData)
      .then(() => {
        toast.info("✏️ Produto atualizado com sucesso!");
        listarProdutos();
        limparFormulario();
      })
      .catch(() => toast.error("❌ Erro ao atualizar produto"));
  };

  const deletarProduto = (id, nome) => {
    if (window.confirm(`Tem certeza que deseja excluir o produto "${nome}"?`)) {
      axios.delete(`http://localhost:8080/produtos/${id}`)
        .then(() => {
          toast.warn("🗑️ Produto excluído com sucesso!");
          listarProdutos();
        })
        .catch(() => toast.error("❌ Erro ao excluir produto"));
    }
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    if (produtoEditando) {
      atualizarProduto();
    } else {
      criarProduto();
    }
  };

  return (
    <>
      <div className="App">
        <header className="header">
          <h1>🛍 Gerenciador de produtos</h1>
          <button
            className="btn btn-primary"
            onClick={() => setMostrarFormulario(!mostrarFormulario)}
          >
            {mostrarFormulario ? "❌ Cancelar" : "➕ Novo Produto"}
          </button>
        </header>

        {/* Formulário */}
        {mostrarFormulario && (
          <div className="formulario-container">
            <h2>{produtoEditando ? "✏ Editar Produto" : "➕ Novo Produto"}</h2>
            <form onSubmit={handleSubmit} className="formulario">
              <div className="form-group">
                <label>Nome:</label>
                <input
                  type="text"
                  name="nome"
                  value={formulario.nome}
                  onChange={handleChange}
                  required
                  placeholder="Ex: Mouse Gamer"
                />
              </div>

              <div className="form-row">
                <div className="form-group">
                  <label>Preço:</label>
                  <input
                    type="number"
                    step="0.01"
                    name="preco"
                    value={formulario.preco}
                    onChange={handleChange}
                    required
                    placeholder="199.90"
                  />
                </div>

                <div className="form-group">
                  <label>Quantidade:</label>
                  <input
                    type="number"
                    name="quantidade"
                    value={formulario.quantidade}
                    onChange={handleChange}
                    required
                    placeholder="10"
                  />
                </div>
              </div>

              <div className="form-group">
                <label>Descrição:</label>
                <textarea
                  name="descricao"
                  value={formulario.descricao}
                  onChange={handleChange}
                  required
                  placeholder="Ex: Mouse com iluminação RGB"
                  rows="3"
                />
              </div>

              <div className="form-actions">
                <button type="submit" className="btn btn-success">
                  {produtoEditando ? "💾 Atualizar" : "➕ Criar"}
                </button>
                <button
                  type="button"
                  className="btn btn-secondary"
                  onClick={limparFormulario}
                >
                  ❌ Cancelar
                </button>
              </div>
            </form>
          </div>
        )}

        {/* Lista de produtos */}
        <div className="produtos-container">
          <h2>📦 Lista de produtos ({produtos.length})</h2>

          {produtos.length === 0 ? (
            <p className="sem-produtos">Nenhum produto cadastrado.</p>
          ) : (
            <div className="produtos-grid">
              {produtos.map(produto => (
                <div key={produto.id} className="produto-card">
                  <div className="produto-header">
                    <h3>{produto.nome}</h3>
                    <span className="produto-id">#{produto.id}</span>
                  </div>

                  <div className="produto-info">
                    <p className="preco">💰 R$ {produto.preco}</p>
                    <p className="quantidade">📦 {produto.quantidade} unidades</p>
                    <p className="descricao">{produto.descricao}</p>
                  </div>

                  <div className="produto-actions">
                    <button
                      className="btn btn-edit"
                      onClick={() => iniciarEdicao(produto)}
                    >
                      ✏ Editar
                    </button>
                    <button
                      className="btn btn-delete"
                      onClick={() => deletarProduto(produto.id, produto.nome)}
                    >
                      🗑 Excluir
                    </button>
                  </div>
                </div>
              ))}
            </div>
          )}
        </div>
      </div>
      <ToastContainer position="top-right" autoClose={3000} />
    </>
  );
}

export default App;

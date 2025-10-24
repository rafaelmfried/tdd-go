package ponteiros_test

import (
	ponteiros "tdd/7-ponteiros"
	"testing"
)

func confirmaErro(t testing.TB, resultado error, esperado string) {
	t.Helper()
	if resultado == nil {
		t.Fatalf("esperava um erro mas não ocorreu nenhum")
	}

	if resultado.Error() != esperado {
		t.Errorf("resultado %q, esperado %q", resultado.Error(), esperado)
	}
}

func confirmaErroInexistente(t testing.TB, resultado error) {
	t.Helper()
	if resultado != nil {
		t.Fatalf("não esperava um erro, mas ocorreu um: %q", resultado)
	}
}

func TestCarteira(t *testing.T) {
		t.Run("Depositar dinheiro na carteira", func(t *testing.T) {
				carteira := ponteiros.NewCarteira(10)
				carteira.Depositar(10)

				resultado := carteira.Saldo()
				esperado := 20

				t.Logf("resultado %d, esperado %d", resultado, esperado)
				if resultado != esperado {
						t.Errorf("resultado %d, esperado %d", resultado, esperado)
				}
		})
}

func TestCarteiraBitcoin(t *testing.T) {
		t.Run("Depositar Bitcoin na carteira", func(t *testing.T) {
				carteira := ponteiros.NewBitcoinCarteira(ponteiros.Bitcoin(20))
				carteira.Depositar(ponteiros.Bitcoin(15))

				resultado := carteira.Saldo()
				esperado := ponteiros.Bitcoin(35)

				t.Logf("resultado %s, esperado %s", resultado, esperado)
				if resultado != esperado {
						t.Errorf("resultado %s, esperado %s", resultado, esperado)
				}
		})
}

func TestCarteirasRetirar(t *testing.T) {
		t.Run("Retirar dinheiro da carteira", func(t *testing.T) {
				carteira := ponteiros.NewCarteira(50)
				erro := carteira.Retirar(20)

				resultado := carteira.Saldo()
				esperado := 30

				t.Logf("resultado %d, esperado %d", resultado, esperado)

				if resultado != esperado {
						t.Errorf("resultado %d, esperado %d", resultado, esperado)
				}

				confirmaErroInexistente(t, erro)
		})

		t.Run("Retirar mais do que o saldo disponível", func(t *testing.T) {
				carteira := ponteiros.NewCarteira(10)
				erro := carteira.Retirar(20)

				resultado := carteira.Saldo()
				esperado := 10

				t.Logf("resultado %d, esperado %d", resultado, esperado)

				confirmaErro(t, erro, ponteiros.ErroSaldoInsuficiente.Error())

				if resultado != esperado || erro == nil {
						t.Errorf("resultado %d, esperado %d", resultado, esperado)
				}
		})

		t.Run("Retirar Bitcoin da carteira", func(t *testing.T) {
				carteira := ponteiros.NewBitcoinCarteira(ponteiros.Bitcoin(50))
				erro := carteira.Retirar(ponteiros.Bitcoin(20))

				resultado := carteira.Saldo()
				esperado := ponteiros.Bitcoin(30)

				t.Logf("resultado %s, esperado %s", resultado, esperado)
				if resultado != esperado {
						t.Errorf("resultado %s, esperado %s", resultado, esperado)
				}
				confirmaErroInexistente(t, erro)
		})

		t.Run("Retirar mais Bitcoin do que o saldo disponível", func(t *testing.T) {
				carteira := ponteiros.NewBitcoinCarteira(ponteiros.Bitcoin(10))
				erro := carteira.Retirar(ponteiros.Bitcoin(20))

				resultado := carteira.Saldo()
				esperado := ponteiros.Bitcoin(10)

				t.Logf("resultado %s, esperado %s", resultado, esperado)

				confirmaErro(t, erro, ponteiros.ErroSaldoInsuficiente.Error())

				if resultado != esperado || erro == nil {
						t.Errorf("resultado %s, esperado %s", resultado, esperado)
				}
		})
}
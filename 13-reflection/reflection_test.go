package reflection_test

import (
	"reflect"
	"testing"
)

type Pessoa struct {
	Nome string
	Endereco Endereco
}

type Endereco struct {
	Cidade string
	Bairro string
	Numero int
}

func obterValor(x interface{}) reflect.Value {
	valor := reflect.ValueOf(x)
	if valor.Kind() == reflect.Ptr {
		valor = valor.Elem()
	}

	return valor
}

func percorre(x interface{}, fn func(string)) {
	valor := obterValor(x)

	percorrerValor := func (valor reflect.Value) {
		percorre(valor.Interface(), fn)
	}

switch valor.Kind() {
	case reflect.String:
		fn(valor.String())
	case reflect.Struct:
		quantidadeDeCampos := valor.NumField()
		for i := 0; i < quantidadeDeCampos; i++ {
			percorrerValor(valor.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < valor.Len(); i++ {
			percorrerValor(valor.Index(i))
		}
	case reflect.Map:
		for _, chave := range valor.MapKeys() {
			percorrerValor(valor.MapIndex(chave))
		}
	}
}

func TestReflection(t *testing.T) {
	t.Run("placeholder", func (t *testing.T) {
		casos := []struct {
			Nome string
			Entrada interface{}
			ChamadasEsperadas []string
		}{
			{
				"Struct com um campo string",
				struct {
					Nome string
				}{"Rafael"},
				[]string{"Rafael"},
			},
			{
				"Struct com dois campos string",
				struct {
					Nome string
					Cidade string
				}{"Rafael", "Salvador"},
				[]string{"Rafael", "Salvador"},
			},
			{
				"Struct com campo não string",
				struct {
					Nome string
					Idade int
				}{"Rafael", 32},
				[]string{"Rafael"},
			},
			{
				"Campos aninhados",
				struct {
					Nome string
					Endereco struct {
						Cidade string
						Bairro string
						Numero int
					}
				}{"Rafael", Endereco{"Salvador", "Pituba", 44}},
				[]string{"Rafael", "Salvador", "Pituba"},
			},
			{
				"Ponteiros para coisas",
				&Pessoa{
					"Rafael",
					Endereco{
						"Salvador",
						"Pituba",
						44,
					},
				},
				[]string{"Rafael", "Salvador", "Pituba"},
			},
			{
				"Slices",
				[]Endereco{
					{"Salvador", "Pituba", 44},
					{"São Paulo", "Moema", 101},
				},
				[]string{"Salvador", "Pituba", "São Paulo", "Moema"},
			},
			{
				"Arrays",
				[2]Endereco{
					{"Salvador", "Pituba", 44},
					{"São Paulo", "Moema", 101},
				},
				[]string{"Salvador", "Pituba", "São Paulo", "Moema"},
			},
			{
				"Maps",
				map[string]string{
					"Um": "Um",
					"Dois": "Dois",
				},
				[]string{"Um", "Dois"},
			},
		}

		for _, caso := range casos {
			t.Run(caso.Nome, func (t *testing.T) {
				var chamadas []string
				percorre(caso.Entrada, func (s string) {
					chamadas = append(chamadas, s)
				})

				if !reflect.DeepEqual(chamadas, caso.ChamadasEsperadas) {
					t.Errorf("Chamadas recebidas %v, esperadas %v", chamadas, caso.ChamadasEsperadas)
				}
			})
		}
	})
}
package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"
)

var test2048Key *rsa.PrivateKey

func init() {
	test2048Key = &rsa.PrivateKey{
		PublicKey: rsa.PublicKey{
			N: fromBase10("14314132931241006650998084889274020608918049032671858325988396851334124245188214251956198731333464217832226406088020736932173064754214329009979944037640912127943488972644697423190955557435910767690712778463524983667852819010259499695177313115447116110358524558307947613422897787329221478860907963827160223559690523660574329011927531289655711860504630573766609239332569210831325633840174683944553667352219670930408593321661375473885147973879086994006440025257225431977751512374815915392249179976902953721486040787792801849818254465486633791826766873076617116727073077821584676715609985777563958286637185868165868520557"),
			E: 3,
		},
		D: fromBase10("9542755287494004433998723259516013739278699355114572217325597900889416163458809501304132487555642811888150937392013824621448709836142886006653296025093941418628992648429798282127303704957273845127141852309016655778568546006839666463451542076964744073572349705538631742281931858219480985907271975884773482372966847639853897890615456605598071088189838676728836833012254065983259638538107719766738032720239892094196108713378822882383694456030043492571063441943847195939549773271694647657549658603365629458610273821292232646334717612674519997533901052790334279661754176490593041941863932308687197618671528035670452762731"),
		Primes: []*big.Int{
			fromBase10("130903255182996722426771613606077755295583329135067340152947172868415809027537376306193179624298874215608270802054347609836776473930072411958753044562214537013874103802006369634761074377213995983876788718033850153719421695468704276694983032644416930879093914927146648402139231293035971427838068945045019075433"),
			fromBase10("109348945610485453577574767652527472924289229538286649661240938988020367005475727988253438647560958573506159449538793540472829815903949343191091817779240101054552748665267574271163617694640513549693841337820602726596756351006149518830932261246698766355347898158548465400674856021497190430791824869615170301029"),
		},
	}
	test2048Key.Precompute()
}

func fromBase10(base10 string) *big.Int {
	i, ok := new(big.Int).SetString(base10, 10)
	if !ok {
		panic("bad number: " + base10)
	}
	return i
}

//  RSA-OAEP加密
func TestEncryptOAEP(t *testing.T) {
	hash := sha256.New()
	random := rand.Reader
	msg := []byte("hello world.")
	label := []byte("orders")

	val, err := rsa.EncryptOAEP(hash, random, &test2048Key.PublicKey, msg, label)
	if err != nil {
		panic(err)
	}
	fmt.Printf("RSA-OAEP encrypt result: %x\n", val)
	// output:
	// RSA-OAEP encrypt result: 66fab38a2be6cd4ffda277b1142db68e4e59c624a9877e0d0f7e91a3c32407c0c63dc5d26e58a6db4a60976eab887d5874e8554e1b5595477b3e7d00bc337f081ae95c6435cf7368e872c3de4997c0cbef08a9cee72a95bbd14fd1bd2a057611e45c929266abf515606235a5ad87e81330428d07b9bc7e9e55002d1614143e8d2763a3762b8909ec803f4fd603d6d43211b19d9deb699e9a9b489e8a47bbc6ed6b107fd88725fce17d67925dc077dc87ae65e3e1357d6cc43e94fc97e3f491be04513220c97d086c1f22d54329bb7786816042e0e4e1f8af888b3709f7e88332af09a89e4191f3de351e4a619e07ddf56bf6ed33721c10692cf418190d6d6d99
}

// RSA_OAEP解密
func TestDecryptOAEP(t *testing.T) {
	// func DecryptOAEP(hash hash.Hash, random io.Reader, priv *PrivateKey, ciphertext []byte, label []byte) ([]byte, error)
	hash := sha256.New()
	random := rand.Reader
	ciphertext, _ := hex.DecodeString("66fab38a2be6cd4ffda277b1142db68e4e59c624a9877e0d0f7e91a3c32407c0c63dc5d26e58a6db4a60976eab887d5874e8554e1b5595477b3e7d00bc337f081ae95c6435cf7368e872c3de4997c0cbef08a9cee72a95bbd14fd1bd2a057611e45c929266abf515606235a5ad87e81330428d07b9bc7e9e55002d1614143e8d2763a3762b8909ec803f4fd603d6d43211b19d9deb699e9a9b489e8a47bbc6ed6b107fd88725fce17d67925dc077dc87ae65e3e1357d6cc43e94fc97e3f491be04513220c97d086c1f22d54329bb7786816042e0e4e1f8af888b3709f7e88332af09a89e4191f3de351e4a619e07ddf56bf6ed33721c10692cf418190d6d6d99")
	label := []byte("orders")

	val, err := rsa.DecryptOAEP(hash, random, test2048Key, ciphertext, label)
	if err != nil {
		panic(err)
	}
	fmt.Println("RSA-OAEP decrypt result:", string(val))
	// output:
	// RSA-OAEP decrypt result: hello world.
}

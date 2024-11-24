# prask-proboj-2024-jesen

# Čo je to Proboj a ako funguje? 

Proboj, skratka pre progamátorský boj, je aktivita z KSP sústredení, kde hráči (vy) programujú vlastného bota, ktorý
súťazí v predom pripravenej hre. K hre je taktiež pripravený template bota, ktorý zvláda komunikáciu so serverom a nejaké
užitočné funkcie. Taktiež obsahuje veľmi jednoduchý príklad jednoduchého bota, ktorého môžete dalej upravovať.

## Štruktúra/harmonogram 

Počas proboja bežia hry (matche), v ktoré sa skladajú z niekoľko stovák kôl, v ktorých vaši boti hrajú. Počas tejto hry 
(matchu) sa nemení mapa, na ktorej hráte a ani ostatní boti, proti ktorým hráte.

Po každej hre (matchy) sa náhodne zvolí mapa a boti, ktorí na nej budú hrať a spustí sa hra (match).

**Začiatok proboju**: XX.11.2024
**Koniec proboju**: XX.11.2024

## Ciele

Zabaviť sa a vyskúsať si niečo pekné nakódiť.

A pre tých kompetetívnejších z vás: Počas hry (matchu) bude váš bot získavať body za rôzne úkony (vid. pravidlá) počas hry. Tieto body sa sčítavajú medzi hrami (matchmi). Kto bude mať na konci najviac bodov, vyhráva.

# Pravidlá hry

## Krátky opis hry

Každý hráč riadi vlastného virtuálneho mafiána. Cieľom hry je zneškodiť všetkých ostatných mafiánov pre tým ako zneškodia
oni vás.

Hra sa hrá na ťahy - váš mafián môže v jednom ťahy vykonať iba jeden úkon, ktorý pridáte do zoznamu úkonov, ktoré
tvoria váš ťah.

## Ako vyzerá mapa

Mapa má tvar kruhu.

Na mate se nachádzajú steny - úsečky, cez ktoré hráč nedokáže prechádzať.

POZOR! Mapa sa časom zmenšuje a výlet za hranicu mapy bude váš mu mafiánovi ubližovať.

## Herné mechaniky

Váš mafián môže v každom ťahu urobiť jeden z týchto úkonov.

### Pohyb

Váš mafián sa môže voľne pohybovať po mape. Môžete sa pohybovať najviac o vzdialenosť *N*.
Pri príkaze na pohyb však možete pouziť ľubovoľnú súradnicu, ak nie je vo vašom dosahu (do vzd. *N*)

### Streľba

Váš mafián môže strielať po protivníkoch. Má rôzne zbrane a tie majú rôzne dostrel. Tiež majú počet nábojov. 
Keď miniete všetky náboje, budete istý počet ťahov prebíjať.

### Prebíjanie

Prebiť môžete taktiež ako samostatný úkon. Kým budete prebíjať nemôžete strielať.

### Zvihnúť zo zeme

Na zemi je kopa rôznych predmentov. Mafián ich môže tieto predmenty zo zeme v ťahu zobrať.

### Vyhodiť zbraň z ruky

Váš mafián môže zahodiť zbraň z ruky.    

## Herné objekty

### Mafián

Mafián je váš hráč. Môžete sa s ním pohybovať, strielať a zbierať (vyhadzovať) predmenty.
Váš mafián začína so 100 bodmi života, bez zbrane a na náhodnej pozícii.
### Zbrane

Zbrane sú predmenty, ktoré sa dajú zozbierať zo zeme. Majú istú kapacitu zásobníku. Keď ju vyčerpáte musite
zbraň prebiť, čo tvrá istý počet ťahov, babygirl.

| Typ        | **Range** | **Damange** | **Reload time** | **Ammo capacity** | 
|------------|-----------|-------------|-----------------|-------------------|
| **None**   | 0         | 0           | 0               | 0                 |
| **Knife**  | 10        | 34          | 0               | 1                 | 
| **Pistol** | 25        | 5           | 2               | 10                | 
| **Tommy**  | 50        | 8           | 4               | 25                |

### Lekárnička

Zdravotnícke oddelienie sústredia zasponzorovalo zopár svojich lečivých zázrakov. Môžete so zeme zdvyhnúť lekárničku,
ktorá vám doplní 75 života.

## Hodnotenie 

Za zabiie protivníka získavate 10 bodov.
Za akékoľvek poškodenie protivníka získavate 1 bod.

# Návod na odovzdávanie
todoo
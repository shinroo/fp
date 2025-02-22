CREATE TABLE IF NOT EXISTS spca (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    latitude DECIMAL(10, 8) NOT NULL,
    longitude DECIMAL(11, 8) NOT NULL,
    website VARCHAR(200),
    address VARCHAR(500) NOT NULL
);

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '47105147-4050-48af-b684-028811220b74',
        'ALBERTON SPCA',
        'General Albert Park 29 Heidelberg Road, Lincoln Road, Alberton, 1449',
        'http://spcaalberton.org.za/',
        -26.2872621109,
        28.1258542996
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '6156d22a-eb70-46da-8b39-e4d39f6983a5',
        'AMAHLATHI SPCA (Stutterheim)',
        'Xholorha, Stutterheim, 4930',
        NULL,
        -32.565722335,
        27.3963212924
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'a94b24b0-b832-425e-9ab0-e06e7f9c0502',
        'AMANZIMTOTI SPCA',
        '1 Nyati Rd, Athlone Park, eManzimtoti, 4126',
        'https://www.spcatoti.co.za/',
        -30.0062690904,
        30.9207469416
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '47ce0d02-8b04-412b-814c-1b4675257ffb',
        'ASSISI SPCA (Humansdorp)',
        'St Francis Dr, Humansdorp, 6300',
        'http://www.spcaassisi.com/',
        -34.0368268693,
        24.7833184943
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'f3875393-1c30-41c2-8d67-362937adcc0d',
        'BALLITO SPCA (DOLPHIN COAST - Branch of DURBAN & COAST SPCA)',
        '22 Sandra Rd, Ballito, 4420',
        NULL,
        -29.543882073,
        31.2144864856
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'a16ba824-b790-47d9-a1b8-6db735373527',
        'BEAUFORT WEST SPCA (KAROO)',
        'cnr Grimbeeck & Jooste Street, Beaufort West, 6970',
        NULL,
        -32.369586,
        22.592221
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'c4afadaf-dfb5-4ab4-8352-149d51d4aa6c',
        'BENONI SPCA',
        '3-2, Klein St, Lakefield, Benoni, 1501',
        NULL,
        -26.178064,
        28.289333
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '6d5979e9-a90d-4159-a4a6-c327dae8c15a',
        'BETHAL SPCA',
        'Jabulani Selepe Street, Bethal, 2309',
        NULL,
        -26.451179,
        29.467527
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '67009a1b-5c01-4e35-9130-566f7b8dd479',
        'BETHLEHEM SPCA',
        '4 Witteberg St, Bethlehem, 9701',
        NULL,
        -28.214165,
        28.299165
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '297094c5-81fa-4b7e-a55d-aa840c8c35c0',
        'BLOEMFONTEIN SPCA',
        'Mcgregor St, White City, Bloemfontein, 9323',
        'https://bloemfonteinspca.co.za/',
        -29.122651,
        26.238812
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'f7ba0c2f-628b-4567-99cd-02722e6a253e',
        'BOKSBURG SPCA',
        'cnr Kruger Street & Railway Street, Boksburg, 1459',
        NULL,
        -26.2161579566,
        28.2630434745
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '156a3330-8856-4428-b5db-c269f000d27e',
        'BRAKPAN SPCA',
        '96 Denne Rd, Witpoortjie 117-Ir, Brakpan, 1543',
        'https://www.brakpanspca.co.za/',
        -26.2808536173,
        28.3515267078
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'a2c421f2-d4a0-4340-94fd-8ff804b07a06',
        'CAPE OF GOOD HOPE SPCA',
        'cnr 1st Ave & First Rd, Grassy Park, Cape Town, 8000',
        'http://www.capespca.co.za/',
        -34.0409504433,
        18.4990058917
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'ea9dd23a-68dd-45b2-a216-4ab01ec27bc3',
        'CARLETONVILLE SPCA',
        'Beerster St, Oberholzer, Carletonville, 2499',
        NULL,
        -26.3395635604,
        27.3915133236
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '04d8bfc1-a631-4b4c-a098-f7ae03e3fbfc',
        'DUNDEE & DISTRICT SPCA',
        '14 Watt Rd, Dundee, 3000',
        NULL,
        -28.1698370907,
        30.2530818638
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'c0054c53-c8aa-4124-a7e4-ba379028de4d',
        'DURBAN & COAST SPCA',
        'Springfield Park, 2 Willowfield Cres, Sea Cow Lake, Durban, 4001',
        'http://www.spcadbn.org.za/',
        -29.8037274694,
        30.9936727045
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '6c1c114d-3f67-43df-98b0-6f2fadf02814',
        'EAST LONDON SPCA',
        'Main Rd, Amalinda, East London, 5247',
        NULL,
        -32.9739003205,
        27.8540788064
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'b97fbcb0-d217-4640-a92d-78a579edcfd9',
        'EDENVALE SPCA',
        'Bhacus Road, Sebenza, Edenvale, 1609',
        'http://www.edenvalespca.co.za/',
        -26.1290927332,
        28.1853008105
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'e30dd8a0-e873-47fb-bca9-88b1469d6e8a',
        'EMPANGENI SPCA',
        'Eshowe St, Empangeni Central, Empangeni, 3880',
        'https://www.empangenispca.co.za/',
        -28.7234140939,
        31.8358675201
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'bb52b7a0-c653-45e7-b04e-036a42225e09',
        'ESHOWE SPCA',
        'Unnamed Road, Eshowe, 3815',
        NULL,
        -28.8904319838,
        31.441484028
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'ac5c71a8-8d3c-492a-b707-fbf822ac6323',
        'FRANSCHHOEK SPCA',
        '6 La Provence St, Le Roux, Franschhoek, 7690',
        'http://spcafhk.co.za/',
        -33.8946038779,
        19.1044207262
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '28c6e325-3089-455c-955f-fdc5e3b64cd9',
        'GARDEN ROUTE SPCA - GEORGE',
        'Ossie Urbanweg, George, 6529',
        'http://www.grspca.co.za/',
        -33.9881155689,
        22.4479290102
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'a9e9e834-ef83-4009-b0c5-548f04440f2a',
        'GARDEN ROUTE SPCA - MOSSEL BAY BRANCH',
        '18 Bill Jeffery St, Boplaas, Mossel Bay, 6506',
        NULL,
        -34.1817878324,
        22.0994767955
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'b3bfa359-a47a-4ccc-926b-0e2805594bfb',
        'GERMISTON & BEDFORDVIEW SPCA',
        '12 Junction Rd, Elandsfontein 90-Ir, Germiston, 1400',
        NULL,
        -26.2082255614,
        28.1529082707
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'dd24ba8a-1cd0-4248-9191-98cbe7cc256a',
        'GRAHAMSTOWN SPCA',
        'Old Bay Road, Makhanda, 6139',
        NULL,
        -33.3193647225,
        26.4982210069
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '934aa76e-d06c-4cdb-ae0c-b650156c171a',
        'GREYTOWN SPCA',
        'Whinstone Farm, 1 R74, Greytown, 3250',
        NULL,
        -29.0494296699,
        30.546301545
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '347024af-6f05-48ec-8b39-9367aead5606',
        'HEIDELBERG SPCA',
        '1 Retief St, Heidelberg, Heidelberg - GP, 1441',
        'http://heidelbergspca.co.za/',
        -26.5009216698,
        28.3703319836
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'deb0cd5d-32fc-4260-8422-4c4028fc23c0',
        'JOHANNESBURG SPCA',
        '5 Benray Rd, Reuven, Johannesburg, 2091',
        'http://www.jhbspca.co.za/',
        -26.2357249829,
        28.0296814593
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '3b0e26a8-4116-4667-b574-8e4542b1850f',
        'KEMPTON PARK SPCA',
        '130 Fitter Rd, Spartan, Kempton Park, 1620',
        'http://www.kemptonspca.co.za/',
        -26.1201024847,
        28.2109046752
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '8d3c256f-e5fb-4fb0-90cf-08edd817b1a0',
        'KIMBERLEY SPCA',
        'Molyneaux Rd, Cassandra, Kimberley, 8301',
        'https://kimberleyspca.com/',
        -28.7497805301,
        24.7955538732
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'f70da52f-9a59-4bfe-affc-24da82151582',
        'KING WILLIAMS TOWN SPCA',
        'Old Stutterheim Road, Qonce, 5601',
        NULL,
        -32.8464388598,
        27.3795021763
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '78a0ca37-4bdb-423f-87d6-20c47b8f67a6',
        'KLOOF & HIGHWAY SPCA',
        '29 Village Rd, Kloof, 3610',
        'http://www.kloofspca.co.za/',
        -29.7978395262,
        30.8296305305
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '65e32758-3700-4451-9180-fb3c5cd72a56',
        'KROONSTAD SPCA',
        'Kraalkop, Maokeng Rural, Maokeng, 9500',
        NULL,
        -27.6644837184,
        27.2561426017
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'c415b95c-5a6d-4678-93d8-e1e0f9d63ee0',
        'LEPHALALE SPCA (ELLISRAS)',
        '1443, Chris Hani Ave, Lephalale, 0555',
        NULL,
        -23.670282764,
        27.7048451881
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '33e5cd09-e03d-4769-b5cb-c6c4d0069019',
        'LETABA SPCA (TZANEEN)',
        '1 Jubilee Rd, Hamaboya, Tzaneen, 2029',
        'http://www.letabaspca-dbv.co.za/',
        -23.8550066854,
        30.1530011138
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '7b6b71c4-4ac4-4342-abef-ab97042d8e0c',
        'LOUIS TRICHARDT SPCA',
        '121 Commercial St, Louis Trichardt, 0922',
        NULL,
        -23.0550343039,
        29.9178328511
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '3b923b5e-994a-4b41-a1fd-ed060e983f82',
        'LOWER SOUTH COAST SPCA',
        'cnr Sea Slopes & Quarry Rd, Uvongo, Margate, 4270',
        'https://lscspca.org.za/',
        -30.8397645861,
        30.3720285319
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'c00fbb80-ec4d-47a0-83b9-d4019f5917d4',
        'MAFIKENG SPCA',
        'Old Cookes lake entrance Moloporivier, Mahikeng, 2735',
        NULL,
        -25.8675444643,
        25.6497234755
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'cf127138-464e-4eb8-9c6e-7249cb52eef5',
        'MIDDELBURG SPCA',
        '41 Wicht St, Middelburg, 1050',
        NULL,
        -25.7786013458,
        29.4809141947
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'b6b96d10-c642-491a-9f76-e15a9046282d',
        'MIDRAND SPCA',
        '6 Dale Rd, Glen Austin AH, Midrand, 1685',
        'http://www.midrandspca.co.za/',
        -25.9813229057,
        28.1648381662
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'ef5bc46e-9c99-46bf-81f3-bf18ce3bef24',
        'MOKOPANE SPCA (POTGIETERSRUS)',
        'Municipal grounds, 1 Zebediela Road, Mokopane, 0601',
        NULL,
        -24.1885320861,
        29.0207915927
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'c31edf76-7290-426e-afac-347275fdcba8',
        'MOOI RIVER & DISTRICT SPCA',
        'Mooi River, 3300',
        NULL,
        -29.2005897645,
        29.9923149816
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'c0cd28d8-5381-4ac6-a57d-9e13c8a826de',
        'MTHATHA SPCA (UMTATA)',
        '1 Eagle St, Umtata Central, Mthatha, 5100',
        NULL,
        -31.5869983246,
        28.8001085673
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '5d08891a-41d7-4323-826c-28ee16fc8225',
        'NELSPRUIT SPCA',
        '6 Meidlinger St, Mbombela, 1200',
        'http://www.spcanelspruit.co.za/',
        -25.4616925261,
        30.9766612578
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '28e58fcb-51c0-4b31-bf63-f76215e22679',
        'NEWCASTLE SPCA',
        'Sampson Avenue, Barry Hertzog Park, Newcastle, 2945',
        NULL,
        -27.7432032014,
        29.946280981
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'b9eea1f5-1848-49f1-bd11-760c91f48cb7',
        'NIGEL SPCA',
        'Greater Nigel, Nigel, 1496',
        NULL,
        -26.3935808737,
        28.4386796487
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'a568eddf-4c89-4df2-95e3-dec542d4d12e',
        'PAARL SPCA',
        '9 Ben Barnard Estate, cnr Old Paarl & Klapmuts - Simondium Rd, Southern Paarl, Paarl, 7646',
        'http://www.paarlspca.co.za/',
        -33.7902556466,
        18.9402115293
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'd6bb5aad-ca3b-4d07-9895-b26b02c1fbd5',
        'PARYS SPCA',
        'Abattoir St, Parys, 9585',
        NULL,
        -26.9076104694,
        27.4654452409
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '51b35c17-2c0d-4c2f-ba31-574632a4889b',
        'PHALABORWA SPCA',
        '1 Copper Rd, Phalaborwa, 1389',
        NULL,
        -23.9512433881,
        31.1557001215
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '21712561-e48f-4f4c-aa6f-9336ce700a24',
        'PIETERMARITZBURG SPCA',
        '235 Woodhouse Rd, Scottsville, Pietermaritzburg, 3201',
        NULL,
        -29.6042077282,
        30.4126045842
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '526fa2f9-4d39-4161-8967-068f1fa87ad5',
        'POLOKWANE SPCA (PIETERSBURG)',
        '46, Lansdale, Roodepoort, Polokwane, 0700',
        NULL,
        -23.9602359628,
        29.4470916056
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '34f70ab5-f9d4-46c0-baba-e20ccf54c2fb',
        'PORT ALFRED & NDLAMBE DISTRICT SPCA',
        '3510 Dickerson Dr, Port Alfred, 6170',
        NULL,
        -33.5831427731,
        26.8899652569
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'a7c6c894-aa83-4fdd-8dc2-512b425128b2',
        'RANDBURG SPCA',
        '229 Northumberland Ave, Northriding, Randburg, 2162',
        'http://www.spca-rbg.org.za/',
        -26.0473896941,
        27.9533050565
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '422fec7b-d235-47f9-b6fb-8e9790ec87e9',
        'RANDWEST SPCA (Randfontein)',
        'Randfontein Rd, Uitvalfontein 244-Iq, Randfontein, 1779',
        NULL,
        -26.1743431289,
        27.7129025694
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'f385c257-cb4c-4593-8c6b-ef5a59771289',
        'RICHARDS BAY SPCA',
        'Bayview Boulevard, Meer En See, Richards Bay, 3901',
        NULL,
        -28.7873469958,
        32.0871644392
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '833506bb-2f21-48e9-afe6-f12b4be04141',
        'ROODEPOORT SPCA',
        'Gauteng',
        'http://www.spcaroodepoort.org.za/',
        -26.1848069259,
        27.9291077198
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '2ae51091-d723-4ad2-8a28-7e3a3fa05c71',
        'RUSTENBURG SPCA',
        '11 Escom St, Rustenburg, 2999',
        NULL,
        -25.6754170049,
        27.2596671447
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '16db9057-21d2-4a32-bb35-8ebd304695ea',
        'SANDTON SPCA',
        '11 9th St, Marlboro Gardens, Marlboro, Sandton, 2146',
        'http://www.sandtonspca.org.za/',
        -26.0943091806,
        28.0878524092
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'e8e1813a-15d7-4d88-8a27-19f88494083e',
        'SPRINGS SPCA',
        'Corner of Ermelo and Townsend Roads, Strubenvale, Springs, 1559',
        NULL,
        -26.2619053889,
        28.4808505963
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '34aa5868-659b-4824-8748-632a163dd140',
        'SWARTLAND & WEST COAST SPCA',
        '551 Yzerfontein Road, Darling, 7345',
        'http://www.swartlandwestcoastspca.co.za/',
        -33.3706558954,
        18.3742580222
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '4ea758c0-6bf1-4e99-b98c-51b97e800c29',
        'SWELLENDAM SPCA',
        '51 Bontebok St, Railton, Swellendam, 6740',
        'http://spcaswellendam.co.za/',
        -34.0385063515,
        20.4415872545
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'e3fe8fca-182f-4563-a072-360f55eaab5a',
        'TSHWANE SPCA (PRETORIA)',
        '316 Petroleum St, Silverton, Pretoria, 0127',
        'http://www.spcapta.org.za/',
        -25.7253156139,
        28.3119561591
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'ccf6573a-9092-40ce-b488-786a0bfb5f41',
        'Centurion',
        'Kruger Ave, Lyttelton Manor, Centurion, 0157',
        NULL,
        -25.8486564012,
        28.2227063381
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '66ae3bbb-efe1-4812-81ab-a35bd9d67d4f',
        'UITENHAGE & DISTRICT SPCA',
        'Schonland St, Cape Road Industrial, Kariega, 6021',
        'http://www.uitenhagespca.co.za/',
        -33.776861356,
        25.3829304799
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'c47378f3-1bd4-4f94-8505-6e794dbb1d78',
        'uMNGENI SPCA (Howick)',
        '15 Campbell Rd, Howick, 3290',
        NULL,
        -29.4913406392,
        30.231921899
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'a3f17a4c-d3ec-49bc-8725-a1e2737617fc',
        'VEREENIGING & VANDERBIJLPARK SPCA',
        '66 General Smuts Rd, Duncanville, Vereeniging, 1939',
        NULL,
        -26.6610342958,
        27.9309335296
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'a83b2d13-482d-4dc1-bf68-02fbce7910d1',
        'VIRGINIA SPCA',
        '9430, Harmony, Virginia, 9430',
        NULL,
        -28.0835089283,
        26.8474192363
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '6ceb9c5f-53db-4f6e-a180-415f63f22f9a',
        'WELLINGTON SPCA',
        'Interpace Street, Wellington, 7655',
        NULL,
        -33.6486930094,
        18.9830343843
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '51e3ecbb-365d-44af-9c0e-81ca73ef5fca',
        'WEST COAST SPCA (Branch of Swartland SPCA)',
        'West Coast Peninsula',
        NULL,
        -32.9222871519,
        18.0493160753
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        'b79d8375-5532-4d70-a008-69f3d5a162f8',
        'WESTONARIA SPCA (Amalgamated with Randfontein SPCA)',
        'Westonaria, 1779',
        NULL,
        -26.3107245072,
        27.6470575385
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '6e86c46e-adea-41a7-9c86-914675ad7fad',
        'WHITE RIVER SPCA',
        '2 Tungsten Rd, White River, 1240',
        NULL,
        -25.3390484542,
        31.0071307252
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '4a877fe8-8c61-4a5b-bd8f-067de1672030',
        'WINELANDS SPCA',
        '537F+5R, Ashton, 6715',
        NULL,
        -33.8371538653,
        20.0745943314
    ) ON CONFLICT (id) DO NOTHING;

INSERT INTO
    spca (id, name, address, website, latitude, longitude)
VALUES
    (
        '8f9f9e13-fcf4-4292-ba94-2ca77c4baf46',
        'WITBANK SPCA',
        '12 Ermelo Ave, River View, eMalahleni, 1034',
        'http://www.witbankspca.org.za/',
        -25.8428296008,
        29.2511136679
    ) ON CONFLICT (id) DO NOTHING;
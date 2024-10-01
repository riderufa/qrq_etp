package api

//func TestAPI_preSearchesHandler(t *testing.T) {
//	dbase := db.New()
//	dbase.NewPreSearch(db.PreSearch{})
//	api := New(dbase)
//	req := httptest.NewRequest(http.MethodGet, "/pre_search/105310", nil)
//	rr := httptest.NewRecorder()
//	api.r.ServeHTTP(rr, req)
//	if !(rr.Code == http.StatusOK) {
//		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
//	}
//	b, err := ioutil.ReadAll(rr.Body)
//	if err != nil {
//		t.Fatalf("не удалось раскодировать ответ сервера: %v", err)
//	}
//	var data []db.PreSearch
//	err = json.Unmarshal(b, &data)
//	if err != nil {
//		t.Fatalf("не удалось раскодировать ответ сервера: %v", err)
//	}
//	const wantLen = 1
//	if len(data) != wantLen {
//		t.Fatalf("получено %d записей, ожидалось %d", len(data), wantLen)
//	}
//}

//func TestAPI_newPreSearchHandler(t *testing.T) {
//	dbase := db.New()
//	dbase.NewPreSearch(db.PreSearch{})
//	api := New(dbase)
//
//	preSearch := db.PreSearch{}
//	b, _ := json.Marshal(preSearch)
//	req := httptest.NewRequest(http.MethodPost, "/pre_search", bytes.NewBuffer(b))
//	rr := httptest.NewRecorder()
//	api.r.ServeHTTP(rr, req)
//	if rr.Code != http.StatusOK {
//		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
//	}
//}

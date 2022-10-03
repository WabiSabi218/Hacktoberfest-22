//Configure SSL (Secure Sockets Layer) certificates for a TCP connection
//to Cloud SQL for PostgreSQL by using Go's database/sql package:

dbRootCert := os.Gentev("DB_ROOT_CERT") //e.g, '/path/to/my/server-ca.pem'
if dbRootCert != "" {
    var (
            dbCert = mustGetenv("DB_CERT") //e.g, '/path/to/my/client-cert.pem'
            dbKey  = mustGetenv("DB_KEY") //e.g, '/path/to/my/client-key.pem'
    )
            dbURI += fmt.Sprintf(" sslmode=require sslrootcert=%s sslcert=%s sslkey=%s", dbRootCert, dbCert, dbKey)
}
package com.fml.blog.user.model;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;
import javax.persistence.Table;

import org.hibernate.annotations.GenericGenerator;

@Entity
@Table(name = "TB_USER")
public class User {
    @Id
    @GenericGenerator(name = "BG_UID", strategy = "uuid")
    @GeneratedValue(generator = "BG_UID")
    @Column(name = "BG_UID", length = 36)
    private String uid;
    
    @Column(name = "BG_USERNAME", length = 36)
    private String username;

    @Column(name = "BG_PASSWORD", length = 36)
    private String password;
    
    @Column(name = "BG_REGISTED", length = 255)
    private String registed;
    
    @Column(name = "BG_LASTLOGIN", length = 36)
    private String lastlogin;
    
    @Column(name = "BG_PORTRAIT", length = 255)
    private String portrait;
    
    public String getUid() {
        return uid;
    }

    public void setUid(String uid) {
        this.uid = uid;
    }

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }
}

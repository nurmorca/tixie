package com.tixie.user.data.entity;

import java.sql.Timestamp;

import jakarta.persistence.*;

@Entity
@Table(name="USERS")
public class User {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name="US_ID")
    private int usId;
    @Column(name="US_FIRST_NAME", nullable = false, length = 80)
    private String usFirstName;
    @Column(name="US_LAST_NAME", nullable = false, length = 80)
    private String usLastName;
    @Column(name="US_EMAIL", nullable = false, unique = true)
    private String usEmail;
    @Column(name="US_PHONE_NUMBER")
    private String usPhoneNumber;
    @Column(name="US_CREATED_AT")
    private Timestamp usCreatedAt;
    @Column(name="US_UPDATED_AT")
    private Timestamp usUpdatedAt;
    @Column(name="US_PASSWORD", nullable = false, length = 200)
    private String usPassword;

    public int getUsId() {
        return usId;
    }

    public void setUsId(int usId) {
        this.usId = usId;
    }

    public String getUsFirstName() {
        return usFirstName;
    }

    public void setUsFirstName(String usFirstName) {
        this.usFirstName = usFirstName;
    }

    public String getUsLastName() {
        return usLastName;
    }

    public void setUsLastName(String usLastName) {
        this.usLastName = usLastName;
    }

    public String getUsEmail() {
        return usEmail;
    }

    public void setUsEmail(String usEmail) {
        this.usEmail = usEmail;
    }

    public String getUsPhoneNumber() {
        return usPhoneNumber;
    }

    public void setUsPhoneNumber(String usPhoneNumber) {
        this.usPhoneNumber = usPhoneNumber;
    }

    public Timestamp getUsCreatedAt() {
        return usCreatedAt;
    }

    public void setUsCreatedAt(Timestamp usCreatedAt) {
        this.usCreatedAt = usCreatedAt;
    }

    public Timestamp getUsUpdatedAt() {
        return usUpdatedAt;
    }

    public void setUsUpdatedAt(Timestamp usUpdatedAt) {
        this.usUpdatedAt = usUpdatedAt;
    }

    public String getUsPassword() {
        return usPassword;
    }

    public void setUsPassword(String usPassword) {
        this.usPassword = usPassword;
    }
}

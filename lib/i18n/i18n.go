/*
Copyright © 2019 Canadian Digital Service <max.neuvians@cds-snc.ca>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package i18n

import (
	"strings"

	"github.com/spf13/viper"
)

type Dict struct {
	en, fr string
}

func Dicts() map[string]Dict {
	m := make(map[string]Dict)

	m["check_root_mfa"] = Dict{
		"Checking AWS root account for MFA ...",
		"Vérification du compte racine AWS pour l'AMF ....",
	}
	m["check_root_mfa_pass"] = Dict{
		"Root MFA is enabled",
		"L'AMF racine est activée",
	}
	m["check_root_mfa_fail"] = Dict{
		"Root MFA is not enabled",
		"L'AMF racine n'est pas activée",
	}

	m["check_root_keys"] = Dict{
		"Checking AWS root account for programmatic keys ...",
		"Vérification du compte racine de la PIC pour les clés programmatiques ....",
	}
	m["check_root_keys_pass"] = Dict{
		"Root has no programmatic keys",
		"La racine n'a pas de clés programmatiques",
	}
	m["check_root_keys_fail"] = Dict{
		"Root has programmatic keys",
		"La racine a des clés programmatiques",
	}

	m["check_user_mfa"] = Dict{
		"Checking AWS console users accounts for MFA ...",
		"Vérification des comptes des utilisateurs de la console AWS pour MFA ...",
	}
	m["check_user_mfa_pass"] = Dict{
		"All user accounts use MFA (taking into account %d breakglass accounts)",
		"Tous les comptes utilisateurs utilisent l'AMF (en tenant compte du %d des comptes breakglass)",
	}
	m["check_user_mfa_fail"] = Dict{
		"%d out of %d console users have MFA active",
		"Les utilisateurs de la console %d sur %d ont MFA actif",
	}

	m["check_admin_users"] = Dict{
		"Checking AWS for users with admin policies attached ...",
		"Vérification de la AWS pour les utilisateurs ayant des politiques d'administration jointes ...",
	}
	m["check_admin_users_pass"] = Dict{
		"No non-accounted for user accounts have admin policies attached",
		"Aucun compte d'utilisateur non comptabilisé n'est associé à des politiques d'administration.",
	}
	m["check_admin_users_fail"] = Dict{
		"%d user(s) have admin policies attached (%d expected)",
		"Le(s) utilisateur(s) %d ont des politiques d'administration jointes (%d attendu)",
	}

	m["check_lambda"] = Dict{
		"Checking AWS for lambda log export function ...",
		"Vérification de la fonction AWS pour l'exportation de billes lambda ...",
	}
	m["check_lambda_pass"] = Dict{
		"Lambda export function found",
		"Fonction d'exportation lambda trouvée",
	}
	m["check_lambda_fail"] = Dict{
		"Lambda export function missing",
		"Fonction d'exportation lambda manquante",
	}

	m["check_password_policy"] = Dict{
		"Checking AWS password policy ...",
		"Vérification de la politique de mot de passe AWS ....",
	}
	m["check_password_policy_pass"] = Dict{
		"Password must be 15 characters or longer",
		"Le mot de passe doit comporter 15 caractères ou plus",
	}
	m["check_password_policy_fail"] = Dict{
		"Password can be less than 15 characters",
		"Le mot de passe peut comporter moins de 15 caractères.",
	}

	m["check_guard_duty"] = Dict{
		"Checking AWS GuardDuty ...",
		"Vérification de AWS GuardDuty ....",
	}
	m["check_guard_duty_pass"] = Dict{
		"GuardDuty found with master account enabled",
		"GuardDuty trouvé avec le compte principal activé",
	}
	m["check_guard_duty_fail"] = Dict{
		"No GuardDuty detectors found!",
		"Aucun détecteur GuardDuty trouvé !",
	}

	m["check_ec2"] = Dict{
		"Checking AWS EC2 data residency ...",
		"Vérification de la résidence des données AWS EC2 ...",
	}
	m["check_ec2_pass"] = Dict{
		"No EC2 instances found outside ca-central-1",
		"Aucun cas d'EC2 trouvé à l'extérieur de ca-central-1",
	}
	m["check_ec2_fail"] = Dict{
		"EC2 instances found outside ca-central-1",
		"EC2 trouvés à l'extérieur de ca-central-1",
	}

	m["check_s3_enc"] = Dict{
		"Checking AWS S3 bucket encryption settings ...",
		"Vérification des paramètres de cryptage du seau AWS S3 ...",
	}
	m["check_s3_enc_pass"] = Dict{
		"No unexpected S3 bucket found without encryption",
		"Aucun seau S3 inattendu trouvé sans cryptage",
	}
	m["check_s3_enc_fail"] = Dict{
		"S3 bucket found without encryption",
		"Seau S3 trouvé sans cryptage",
	}

	m["check_rds_enc"] = Dict{
		"Checking AWS RDS encryption settings ...",
		"Vérification des paramètres de cryptage AWS RDS ...",
	}
	m["check_rds_enc_pass"] = Dict{
		"No RDS instance found without encryption",
		"Aucune instance RDS trouvée sans cryptage",
	}
	m["check_rds_enc_fail"] = Dict{
		"RDS instance found without encryption",
		"Instance RDS trouvée sans cryptage",
	}

	m["check_sg_p80"] = Dict{
		"Checking AWS EC2 security groups for port 80 ingress ...",
		"Vérification des groupes de sécurité AWS EC2 pour l'entrée du port 80 ....",
	}
	m["check_sg_p80_pass"] = Dict{
		"No security groups found exposing port 80",
		"Aucun groupe de sécurité n'a trouvé le port d'exposition 80",
	}
	m["check_sg_p80_fail"] = Dict{
		"Security group with port 80 found",
		"Groupe de sécurité avec port 80 trouvé",
	}

	return m
}

func T(key string) string {
	dict := Dicts()

	if _, ok := dict[key]; !ok {
		return key
	}

	if strings.Contains(viper.GetString("lang"), "fr") {
		return dict[key].fr
	} else {
		return dict[key].en
	}
}
